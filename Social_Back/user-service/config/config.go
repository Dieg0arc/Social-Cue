package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func InitMongo() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando .env")
	}

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error conectando a MongoDB:", err)
	}

	DB = client.Database(os.Getenv("MONGODB_NAME"))
	log.Println("MongoDB conectado")
}

func GetCollection(name string) *mongo.Collection {
	return DB.Collection(name)
}
