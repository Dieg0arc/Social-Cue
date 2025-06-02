package handlers

import (
    "net/http"
    
    "github.com/labstack/echo/v4"
)

// HealthCheck verifica que el servidor esté funcionando
func HealthCheck(c echo.Context) error {
    return c.JSON(http.StatusOK, map[string]string{
        "status": "ok",
    })
}