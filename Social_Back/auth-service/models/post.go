package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    AuthorID  primitive.ObjectID `bson:"author_id" json:"author_id"`
    Content   string             `bson:"content" json:"content"`
    Likes     int                `bson:"likes" json:"likes"`
    CreatedAt int64              `bson:"created_at" json:"created_at"`
}
