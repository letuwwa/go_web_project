package handlers

import (
	"github.com/labstack/echo/v4"
	"go_web_project/config"
	"go_web_project/models"
	"go_web_project/utils"
	"net/http"
)

func GetUsers(c echo.Context) error {
	var records []models.User
	result := config.DBInit().Find(&records)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{Status: utils.DBOperationError})
	}
	return c.JSON(http.StatusOK, records)
}

func GetUserByID(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	result := config.DBInit().First(&user, "id = ?", id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, utils.Response{Status: utils.DBEntityNotFund})
	}
	return c.JSON(http.StatusOK, user)
}

func DeleteUserByID(c echo.Context) error {
	id := c.Param("id")
	result := config.DBInit().Delete(&models.User{}, "id = ?", id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.DBOperationError})
	}
	return c.JSON(http.StatusOK, utils.Response{Status: utils.DBOperationSuccess})
}
