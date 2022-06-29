package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, "Go Web Project")
}
