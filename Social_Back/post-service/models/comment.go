package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Comentario representa un comentario hecho en un post.
type Comentario struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PostID    primitive.ObjectID `bson:"postId,omitempty" json:"postId"`
	AutorID   primitive.ObjectID `bson:"autorId,omitempty" json:"autorId"`
	Contenido string             `bson:"contenido,omitempty" json:"contenido"`
	Fecha     time.Time          `bson:"fecha,omitempty" json:"fecha"`
}
