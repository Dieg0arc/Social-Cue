package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Notificacion representa una notificación enviada a un usuario.
type Notificacion struct {
	ID      primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	Tipo    string              `bson:"tipo" json:"tipo"` // Ej: "comentario", "like"
	De      primitive.ObjectID  `bson:"de" json:"de"`     // Quién la genera
	Para    primitive.ObjectID  `bson:"para" json:"para"` // Receptor
	Mensaje string              `bson:"mensaje" json:"mensaje"`
	Leida   bool                `bson:"leida" json:"leida"`
	Fecha   time.Time           `bson:"fecha" json:"fecha"`
	PostID  *primitive.ObjectID `bson:"postId,omitempty" json:"postId"`
}
