package handlers

import (
	"github.com/labstack/echo/v4"
	"go_web_project/config"
	"go_web_project/models"
	"go_web_project/utils"
	"net/http"
)

func GetUserByID(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	result := config.DBInit().First(&user, "id = ?", id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, utils.Response{Status: utils.DBEntityNotFund})
	}
	return c.JSON(http.StatusOK, user)
}
