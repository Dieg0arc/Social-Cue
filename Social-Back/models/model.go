package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User define el modelo de usuario
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string             `bson:"username" json:"username"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"-"` // No exponer en JSON
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

// LoginRequest define la estructura para solicitudes de inicio de sesión
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterRequest define la estructura para solicitudes de registro
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse estructura para la respuesta de login
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// RegisterResponse estructura para la respuesta de registro
type RegisterResponse struct {
	Message string `json:"message"`
	UserID  string `json:"user_id"`
}
