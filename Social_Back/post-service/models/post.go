package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Titulo      string             `bson:"titulo" json:"titulo"`
	Contenido   string             `bson:"contenido" json:"contenido"`
	Tipo        string             `bson:"tipo" json:"tipo"` // video, tutorial, documento
	Categoria   string             `bson:"categoria" json:"categoria"`
	Tags        []string           `bson:"tags" json:"tags"`
	URLArchivo  string             `bson:"urlArchivo,omitempty" json:"urlArchivo,omitempty"` // si se sube
	AutorID     primitive.ObjectID `bson:"autorId" json:"autorId"`
	FechaCreado time.Time          `bson:"fechaCreado" json:"fechaCreado"`
}
