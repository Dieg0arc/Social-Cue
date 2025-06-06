package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret-key") // ⚠️ En producción, usar variable de entorno

// Claims define los datos que se incluirán en el token
type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Rol    string `json:"rol"`
	jwt.RegisteredClaims
}

// GenerateJWT genera un token válido para el usuario
func GenerateJWT(userID, email, rol string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		Email:  email,
		Rol:    rol,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// GetJWTKey devuelve la clave secreta para firmar/verificar el token
func GetJWTKey() []byte {
	return jwtKey
}
