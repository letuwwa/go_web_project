package main

import (
	"github.com/labstack/echo/v4"
	"go_web_project/config"
	"go_web_project/handlers"
)

func main() {
	config.DBInit()

	e := echo.New()
	e.GET("/", handlers.Index)
	e.POST("/user", handlers.AddUser)

	e.Logger.Fatal(e.Start(":1323"))
}
