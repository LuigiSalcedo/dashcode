package controllers

import (
	"dashcode/models"
	"dashcode/security"
	"dashcode/services"
	"dashcode/services/login"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Endpoint to login -> /login
func Login(c echo.Context) error {
	loginData := models.LoginModel{}
	err := json.NewDecoder(c.Request().Body).Decode(&loginData)

	if err != nil {
		return c.JSON(http.StatusBadRequest, services.ErrorBadRequest)
	}

	id, srvErr := login.Login(loginData.Email, loginData.Password)

	if srvErr != nil {
		return c.JSON(srvErr.Code, srvErr)
	}

	return c.JSON(http.StatusOK, security.GenerateJWT(id))
}
