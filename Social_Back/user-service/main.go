package main

import (
	"log"
	"os"

	"user-service/config"
	"user-service/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	config.InitMongo()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	routes.RegisterRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	log.Println("user-service corriendo en puerto " + port)
	e.Logger.Fatal(e.Start(":" + port))
}
