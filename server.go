package main

import (
	"github.com/labstack/echo/v4"
	"go_web_project/config"
	"go_web_project/handlers"
	"go_web_project/middleware"
	"go_web_project/models"
)

func main() {
	db := config.DBInit()
	db.AutoMigrate(&models.User{})

	e := echo.New()
	e.GET("/", handlers.Index)
	e.GET("/user/:id", handlers.GetUserByID, middleware.AuthMiddleware)

	e.POST("/auth/login", handlers.AuthLogin)
	e.POST("/auth/register", handlers.AuthRegister)

	e.Logger.Fatal(e.Start(":1323"))
}
