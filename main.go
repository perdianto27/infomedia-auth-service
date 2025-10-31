package main

import (
	"main/configs"
	"main/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// loadEnv()
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoutes(e)
	e.Start(":" + os.Getenv("PORT"))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
