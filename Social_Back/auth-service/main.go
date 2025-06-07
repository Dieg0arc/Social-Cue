package main

import (
	"log"

	"authservice/config"
	"authservice/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Carga variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		log.Println("Advertencia: No se pudo cargar el archivo .env")
	}

	// Cargar configuración y conectar a MongoDB
	cfg := config.GetConfig()
	config.InitMongo(cfg)

	// Inicializar Echo
	e := echo.New()

	// Middleware de logging y recuperación de errores
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Middleware CORS para permitir llamadas desde el frontend
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// Definir rutas
	routes.SetupRoutes(e, cfg)

	// Iniciar servidor
	log.Fatal(e.Start(":8080"))
}
