package handlers

import (
	"github.com/labstack/echo/v4"
	"go_web_project/config"
	"net/http"
)

func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, config.Response{Status: "index page"})
}
