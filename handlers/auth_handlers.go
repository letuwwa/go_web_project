package handlers

import (
	"github.com/labstack/echo/v4"
	"go_web_project/config"
	"go_web_project/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type RequestUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthRegister(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	var hashPassword, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, config.Response{Status: config.HashOperationError})
	}
	user.Password = string(hashPassword)
	result := config.DBInit().Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, config.Response{Status: config.DBOperationError})
	}
	user.Password = "hash"
	return c.JSON(http.StatusOK, config.Response{Status: config.DBOperationSuccess})
}

func AuthLogin(c echo.Context) error {
	requestUser := new(RequestUser)
	if err := c.Bind(requestUser); err != nil {
		return c.JSON(http.StatusBadRequest, config.Response{Status: err.Error()})
	}
	var DBUser models.User
	result := config.DBInit().Where("username = ?", requestUser.Username).First(&DBUser)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, config.Response{Status: config.DBEntityNotFund})
	}
	err := bcrypt.CompareHashAndPassword([]byte(DBUser.Password), []byte(requestUser.Password))
	if err != nil {
		return c.JSON(http.StatusForbidden, config.Response{Status: config.LoginError})
	}
	return c.JSON(http.StatusOK, config.Response{Status: config.LoginSuccess})
}
