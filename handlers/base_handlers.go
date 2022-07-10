package handlers

import (
	"github.com/labstack/echo/v4"
	"go_web_project/utils"
	"net/http"
)

func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, utils.Response{Status: "index page"})
}
