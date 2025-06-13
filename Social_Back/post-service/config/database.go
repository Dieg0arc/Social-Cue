package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// ConnectDB establece la conexión con MongoDB
func ConnectDB() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		panic(fmt.Sprintf("Error creando el cliente de MongoDB: %v", err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		panic(fmt.Sprintf("Error conectando a MongoDB: %v", err))
	}

	// Selecciona la base de datos
	DB = client.Database("social_cue")
	fmt.Println(" Conexión a MongoDB exitosa")
}

// GetCollection obtiene una colección específica
func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}
