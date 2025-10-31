package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	usercontroller "main/controllers/user"
	authController "main/controllers/auth"
	middlewares "main/middlewares"
)

func InitRoutes(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.POST("/register", usercontroller.AddUserController)
	e.POST("/login", authController.LoginController)
	e.POST("/refresh", authController.RefreshTokenController)
	e.GET("/me", usercontroller.MeController, middlewares.AuthMiddleware)
}
