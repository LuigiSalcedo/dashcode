package controllers

import (
	"dashcode/security"
	"dashcode/services"

	"github.com/labstack/echo/v4"
)

func getIdFromJWT(c echo.Context) (int64, *services.Error) {
	jwt := c.Request().Header.Get("Authorization")

	id, err := security.IDFromJWT(jwt)

	if err != nil {
		return -1, services.ErrorJWT
	}

	return id, nil
}
