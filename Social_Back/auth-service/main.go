package main

import (
	"log"

	"authservice/config"
	"authservice/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	config.InitMongo(cfg)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	routes.SetupRoutes(e, cfg) // ✅ PASA cfg AQUÍ

	log.Fatal(e.Start(":8080"))
}
