package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	MongoURI string
	DBName   string
}

var MongoClient *mongo.Client

func GetConfig() *Config {
	return &Config{
		MongoURI: "mongodb://localhost:27017",
		DBName:   "App_web",
	}
}

func InitMongo(cfg *Config) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatal("Error conectando a MongoDB:", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Error haciendo ping a MongoDB:", err)
	}

	MongoClient = client
	log.Println("MongoDB conectado exitosamente")
}

func GetCollection(name string, cfg *Config) *mongo.Collection {
	return MongoClient.Database(cfg.DBName).Collection(name)
}
