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
	e.GET("/groups/owner", controllers.FetchGroupsByOwner)
	e.GET("/groups/member", controllers.FetchGroupsByMember)
	e.POST("/groups/invite/:groupId", controllers.SendInvitations)

	log.Fatal(e.Start(":8080"))
}
