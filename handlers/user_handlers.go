package handlers

import (
	"github.com/labstack/echo/v4"
	db "go_web_project/config"
	"go_web_project/models"
	"net/http"
)

func AddUser(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	result := db.Init().Create(&u)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, "db operation fail")
	}
	return c.JSON(http.StatusOK, "ok")
}
