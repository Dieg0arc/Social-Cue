package routes

import (
	"post-service/ws"

	"github.com/labstack/echo/v4"
)

func WebSocketRoutes(e *echo.Echo) {
	e.GET("/ws/:id", ws.WebSocketHandler)
}
