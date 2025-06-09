package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Like struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PostID    primitive.ObjectID `bson:"postId" json:"postId"`
	UsuarioID primitive.ObjectID `bson:"usuarioId" json:"usuarioId"`
	Fecha     time.Time          `bson:"fecha" json:"fecha"`
}
