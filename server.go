package main

import (
	"dashcode/controllers"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/login", controllers.Login)
	e.POST("/register", controllers.SaveUser)
	e.POST("/groups", controllers.CreateGroup)
	e.GET("/groups/:id/owner", controllers.FetchGroupsByOwner)

	log.Fatal(e.Start(":8080"))
}
