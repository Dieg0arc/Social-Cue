package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PasswordReset representa un token temporal para restablecer contraseña.
// El código se almacena en hash para mayor seguridad.
type PasswordReset struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	CodeHash  string             `bson:"code_hash"`
	ExpiresAt time.Time          `bson:"expires_at"`
}
