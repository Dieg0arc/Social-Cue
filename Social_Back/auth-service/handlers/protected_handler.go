package handlers

import (
	"net/http"

	"authservice/utils"

	"github.com/labstack/echo/v4"
)

func ProfileHandler(c echo.Context) error {
	claims, ok := c.Get("user").(*utils.Claims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "No autorizado"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user_id": claims.UserID,
		"email":   claims.Email,
		"rol":     claims.Rol,
	})
}
