package controllers

import (
	"dashcode/models"
	"dashcode/services"
	"dashcode/services/users"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Endpoint to save user - /users/register -> POST
func SaveUser(c echo.Context) error {
	u := models.FullUserModel{}

	err := json.NewDecoder(c.Request().Body).Decode(&u)

	if err != nil {
		return c.JSON(http.StatusBadRequest, services.ErrBadRequest)
	}

	srvErr := users.SaveUser(u.Id, u.Name, u.Email, u.Password)

	if err != nil {
		return c.JSON(srvErr.Code, srvErr)
	}

	return c.NoContent(http.StatusCreated)
}