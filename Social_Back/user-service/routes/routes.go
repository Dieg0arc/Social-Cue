package routes

import (
	"user-service/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/users")
	api.GET("/:id", handlers.GetUserByID)
}
