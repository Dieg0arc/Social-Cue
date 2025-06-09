package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Follow struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SeguidorID primitive.ObjectID `bson:"seguidorId" json:"seguidorId"`
	SeguidoID  primitive.ObjectID `bson:"seguidoId" json:"seguidoId"`
	Fecha      time.Time          `bson:"fecha" json:"fecha"`
}
