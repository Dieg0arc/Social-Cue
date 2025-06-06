package routes

import (
	"authservice/config"
	"authservice/handlers"
	"authservice/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, cfg *config.Config) {
	auth := e.Group("/api/auth")
	auth.POST("/login", handlers.LoginHandler(cfg))
	auth.POST("/register", handlers.RegisterHandler(cfg))
	auth.GET("/profile", handlers.ProfileHandler, middleware.JWTMiddleware)
}
