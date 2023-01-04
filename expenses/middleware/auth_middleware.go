package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type response struct {
	Message string `json:"message"`
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")

		if auth != "November 10, 2009" {
			return c.JSON(http.StatusUnauthorized, response{
				Message: "[Error] your authorazation key invalid!",
			})
		}
		return next(c)
	}
}
