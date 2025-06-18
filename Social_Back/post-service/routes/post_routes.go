package routes

import (
	"post-service/controllers"

	"github.com/labstack/echo/v4"
)

func PostRoutes(e *echo.Echo) {
	// Ruta de prueba
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	// CRUD de posts
	e.POST("/posts", controllers.CrearPost)   // Crear post
	e.GET("/posts", controllers.ObtenerPosts) // Obtener todos (con filtros)
	// e.GET("/posts/:id", controllers.ObtenerPostPorID) // Obtener uno por ID
	e.GET("/posts/usuario/:id", controllers.ObtenerPostsPorUsuario)
	e.DELETE("/posts/:id", controllers.EliminarPost) // Eliminar post
	// Comentarios
	e.POST("/posts/:id/comments", controllers.CrearComentario)
	e.GET("/posts/:id/comments", controllers.ObtenerComentarios)
	// Likes
	e.POST("/posts/:id/like", controllers.ToggleLike)
}
