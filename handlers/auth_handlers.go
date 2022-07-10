package handlers

import (
	"github.com/labstack/echo/v4"
	"go_web_project/config"
	"go_web_project/models"
	"go_web_project/utils"
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
		return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.HashOperationError})
	}
	user.Password = string(hashPassword)
	result := config.DBInit().Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{Status: utils.DBOperationError})
	}
	user.Password = "hash"
	return c.JSON(http.StatusOK, utils.Response{Status: utils.DBOperationSuccess})
}

func AuthLogin(c echo.Context) error {
	requestUser := new(RequestUser)
	if err := c.Bind(requestUser); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{Status: err.Error()})
	}
	var DBUser models.User
	result := config.DBInit().Where("username = ?", requestUser.Username).First(&DBUser)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, utils.Response{Status: utils.DBEntityNotFund})
	}
	err := bcrypt.CompareHashAndPassword([]byte(DBUser.Password), []byte(requestUser.Password))
	if err != nil {
		return c.JSON(http.StatusForbidden, utils.Response{Status: utils.LoginError})
	}
	return c.JSON(http.StatusOK, utils.Response{Status: utils.LoginSuccess})
}
