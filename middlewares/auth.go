package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"main/helpers"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "Authorization header missing",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := helpers.VerifyAccessToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "Invalid or expired token",
			})
		}

		c.Set("user", claims) // simpan data user di context
		return next(c)
	}
}