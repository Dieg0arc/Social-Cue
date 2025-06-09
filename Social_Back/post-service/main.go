package main

import (
	"log"
	"post-service/config"
	"post-service/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar .env, se usará configuración por defecto")
	}

	// Conectar a la base de datos
	config.ConnectDB()

	// Inicializar Echo
	e := echo.New()

	// Cargar rutas
	routes.PostRoutes(e)
	routes.UserRoutes(e)
	routes.WebSocketRoutes(e)

	// Arrancar servidor
	e.Logger.Fatal(e.Start(":8081")) // Puerto del microservicio post
}
