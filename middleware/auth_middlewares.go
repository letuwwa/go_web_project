package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func TestMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("middleware is active")
		return next(c)
	}
}
