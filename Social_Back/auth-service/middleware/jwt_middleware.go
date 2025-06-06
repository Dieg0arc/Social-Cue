package middleware

import (
	"net/http"
	"strings"

	"authservice/utils"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Falta el token de autorización"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Formato de token inválido"})
		}

		tokenStr := parts[1]
		claims := &utils.Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.GetJWTKey()), nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Token inválido o expirado"})
		}

		c.Set("user", claims)
		return next(c)
	}
}
