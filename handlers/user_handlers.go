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
		return c.JSON(http.StatusBadRequest, config.Response{Message: config.DBOperationError})
	}
	return c.JSON(http.StatusOK, user)
}

func GetUsers(c echo.Context) error {
	var records []models.User
	result := config.DBInit().Find(&records)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, config.Response{Message: config.DBOperationError})
	}
	return c.JSON(http.StatusOK, records)
}

func GetUserByID(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	result := config.DBInit().First(&user, "id = ?", id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, config.Response{Message: config.DBEntityNotFund})
	}
	return c.JSON(http.StatusOK, user)
}

func DeleteUserByID(c echo.Context) error {
	id := c.Param("id")
	result := config.DBInit().Delete(&models.User{}, "id = ?", id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, config.ResponseDeleted{IsDeleted: false})
	}
	return c.JSON(http.StatusOK, config.ResponseDeleted{IsDeleted: true})
}
