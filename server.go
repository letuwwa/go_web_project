package main

import (
	"github.com/labstack/echo/v4"
	"go_web_project/config"
	"go_web_project/handlers"
	"go_web_project/models"
)

func main() {
	db := config.DBInit()
	db.AutoMigrate(&models.User{})

	e := echo.New()
	e.GET("/", handlers.Index)
	e.GET("/user", handlers.GetUsers)
	e.POST("/user", handlers.AddUser)
	e.GET("/user/:id", handlers.GetUserByID)
	e.DELETE("/user/:id", handlers.DeleteUserByID)

	e.Logger.Fatal(e.Start(":1323"))
}
