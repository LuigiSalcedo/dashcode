package controllers

import (
	"dashcode/models"
	"dashcode/security"
	"dashcode/services"
	"dashcode/services/groups"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Create a group in database -> /groups -> POST
func CreateGroup(c echo.Context) error {
	jwt := c.Request().Header.Get("Authorization")

	if len(jwt) == 0 {
		return c.JSON(services.ErrorJWT.Code, services.ErrorJWT)
	}

	id, err := security.IDFromJWT(jwt)

	if err != nil {
		return c.JSON(services.ErrorJWT.Code, services.ErrorJWT)
	}

	if id == -1 {
		return c.JSON(services.ErrorInternal.Code, services.ErrorInternal)
	}

	g := models.CreateGroup{}

	err = json.NewDecoder(c.Request().Body).Decode(&g)

	if err != nil {
		return c.JSON(services.ErrorJson.Code, services.ErrorJson)
	}

	srvErr := groups.CreateGroup(id, g.Name, g.Description)

	if srvErr != nil {
		return c.JSON(srvErr.Code, srvErr)
	}

	return c.NoContent(http.StatusCreated)
}

// Get groups where the user is the owner /groups/:id/owner
func FetchGroupsByOwner(c echo.Context) error {
	jwt := c.Request().Header.Get("Authorization")

	if len(jwt) == 0 {
		return c.JSON(services.ErrorJWT.Code, services.ErrorJWT)
	}

	id, err := security.IDFromJWT(jwt)

	if err != nil {
		return c.JSON(services.ErrorJWT.Code, services.ErrorJWT)
	}

	idParam := c.Param("id")

	idParamValue, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(services.ErrorPathParam.Code, services.ErrorPathParam)
	}

	if int64(idParamValue) != id {
		return c.JSON(services.ErrorUnauthorized.Code, services.ErrorUnauthorized)
	}

	if err != nil {
		return c.JSON(services.ErrorJson.Code, services.ErrorJson)
	}

	g, srvErr := groups.FetchGroupsByOwner(id)

	if srvErr != nil {
		return c.JSON(srvErr.Code, srvErr)
	}

	return c.JSON(http.StatusOK, g)
}

// Get groups where the user with the id is a member
func FetchGroupsByMember(c echo.Context) error {
	jwt := c.Request().Header.Get("Authorization")

	if len(jwt) == 0 {
		return c.JSON(services.ErrorJWT.Code, services.ErrorJWT)
	}

	id, err := security.IDFromJWT(jwt)

	if err != nil {
		return c.JSON(services.ErrorJWT.Code, services.ErrorJWT)
	}

	idParam := c.Param("id")

	idValue, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(services.ErrorPathParam.Code, services.ErrorPathParam)
	}

	if int64(idValue) != id {
		return c.JSON(services.ErrorUnauthorized.Code, services.ErrorUnauthorized)
	}

	g, srvErr := groups.FetchByMember(id)

	if srvErr != nil {
		return c.JSON(srvErr.Code, srvErr)
	}

	return c.JSON(http.StatusOK, g)
}
