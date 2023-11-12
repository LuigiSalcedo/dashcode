package controllers

import (
	"dashcode/security"
	"dashcode/services"
	"dashcode/services/invitations"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FetchInvitations(c echo.Context) error {
	jwt := c.Request().Header.Get("Authorization")

	if len(jwt) == 0 {
		return c.JSON(services.ErrorJWT.Code, services.ErrorJWT)
	}

	id, err := security.IDFromJWT(jwt)

	if err != nil {
		return c.JSON(services.ErrorJWT.Code, services.ErrorJWT)
	}

	invs, srvErr := invitations.FetchInvitations(id)

	if srvErr != nil {
		return c.JSON(srvErr.Code, srvErr)
	}

	return c.JSON(http.StatusOK, invs)
}
