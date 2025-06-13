package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User representa un documento de usuario en la colección "users"
type User struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username            string             `bson:"username" json:"username"`
	Email               string             `bson:"email" json:"email"`
	Rol                 string             `bson:"rol,omitempty" json:"rol,omitempty"`
	CodigoInstitucional string             `bson:"codigo_institucional,omitempty" json:"codigo_institucional,omitempty"`
}
