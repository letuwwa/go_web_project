package main

import (
	"github.com/labstack/echo/v4"
	"go_web_project/config"
	"go_web_project/handlers"
)

func main() {
	db.Init()

	e := echo.New()
	e.GET("/", handlers.Hello)

	e.Logger.Fatal(e.Start(":1323"))
}
