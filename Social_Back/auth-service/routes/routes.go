package routes

import (
        "authservice/config"
        "authservice/handlers"
        "authservice/middleware"
        "authservice/realtime"

        "github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, cfg *config.Config, hub *realtime.Hub) {
        auth := e.Group("/api/auth")
        auth.POST("/login", handlers.LoginHandler(cfg))
        auth.POST("/register", handlers.RegisterHandler(cfg))
        auth.GET("/profile", handlers.ProfileHandler, middleware.JWTMiddleware)

        posts := e.Group("/api/posts")
        posts.POST("/:id/like", handlers.LikePostHandler(cfg, hub))

        e.GET("/ws", echo.WrapHandler(handlers.WebSocketHandler(hub)))
}
