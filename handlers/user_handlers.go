package handlers

import (
	"github.com/labstack/echo/v4"
	"go_web_project/config"
	"go_web_project/models"
	"net/http"
)

func AddUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	result := config.DBInit().Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, "db operation fail")
	}
	return c.JSON(http.StatusOK, user)
}
