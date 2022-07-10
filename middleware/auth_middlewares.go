package middleware

import (
	"github.com/labstack/echo/v4"
	"go_web_project/utils"
	"net/http"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		JWTToken := c.Request().Header["Token"]
		if len(JWTToken) == 0 {
			return c.JSON(http.StatusForbidden, utils.Response{Status: utils.JWTOperationError})
		}
		token, err := utils.ValidateJWT(JWTToken[0])
		if err != nil {
			return c.JSON(http.StatusForbidden, utils.Response{Status: err.Error()})
		}
		if token.Valid == false {
			return c.JSON(http.StatusForbidden, utils.Response{Status: err.Error()})
		}
		return next(c)
	}
}
