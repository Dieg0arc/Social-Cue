package routes

import (
	"backend-instagram/handlers"
	"github.com/labstack/echo/v4"
)

// SetupRoutes configura todas las rutas de la API
func SetupRoutes(e *echo.Echo) {
	// Grupo de rutas API
	api := e.Group("/api")

	// Ruta de prueba
	api.GET("/health", handlers.HealthCheck)

	// Rutas de autenticación
	auth := api.Group("/auth")
	auth.POST("/login", handlers.Login)
	auth.POST("/register", handlers.Register)
}
