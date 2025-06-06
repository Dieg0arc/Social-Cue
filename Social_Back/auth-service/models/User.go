package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User representa un usuario en la base de datos
type User struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username            string             `bson:"username" json:"username"`
	Email               string             `bson:"email" json:"email"`
	Password            string             `bson:"password" json:"-"`
	Rol                 string             `bson:"rol" json:"rol"`
	CodigoInstitucional string             `bson:"codigo_institucional" json:"codigoInstitucional"`
	CreatedAt           time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt           time.Time          `bson:"updated_at" json:"updated_at"`
}

// LoginRequest para solicitudes de inicio de sesión
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterRequest para solicitudes de registro
type RegisterRequest struct {
	Username            string `json:"username"`
	Email               string `json:"email"`
	Password            string `json:"password"`
	Rol                 string `json:"rol"`
	CodigoInstitucional string `json:"codigoInstitucional"`
}

// LoginResponse para respuestas de inicio de sesión
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// RegisterResponse para respuestas de registro
type RegisterResponse struct {
	Message string `json:"message"`
	UserID  string `json:"user_id"`
}
