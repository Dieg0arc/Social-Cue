package routes

import (
	"post-service/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	e.POST("/users/:id/follow", controllers.SeguirUsuario)
	e.DELETE("/users/:id/unfollow", controllers.DejarDeSeguir)
	e.GET("/users/:id/following", controllers.VerSeguidos)
	e.GET("/users/:id/followers", controllers.VerSeguidores)
	e.GET("/users/:id/notificaciones", controllers.VerNotificaciones)
	e.PATCH("/users/:id/notificaciones/:notiId/leida", controllers.MarcarNotificacionLeida)
}
