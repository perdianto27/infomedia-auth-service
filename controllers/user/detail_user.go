package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func MeController(c echo.Context) error {
	userData := c.Get("user")
	if userData == nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"status":  false,
			"message": "Unauthorized",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  true,
		"message": "Success",
		"data":    userData,
	})
}