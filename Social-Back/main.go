package main

import (
    "log"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "backend-instagram/handlers"
    "backend-instagram/routes"
)

func main() {
    // Inicializar MongoDB
    handlers.InitMongo()

    // Inicializar Echo
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Configurar CORS
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://localhost:3000"}, // URL del frontend Nuxt
        AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
    }))

    // Configurar rutas
    routes.SetupRoutes(e)

    // Iniciar servidor
    log.Fatal(e.Start(":8080"))
}
