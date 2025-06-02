package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config estructura para la configuración
type Config struct {
	MongoURI        string
	DBName          string
	CollectionUsers string
}

// GetConfig devuelve la configuración de la aplicación
func GetConfig() *Config {
	return &Config{
		MongoURI:        "mongodb://localhost:27017",
		DBName:          "App_web",
		CollectionUsers: "users",
	}
}

// GetMongoClient establece conexión con MongoDB
func GetMongoClient(cfg *Config) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatal(err)
	}

	// Verificar conexión
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Conexión exitosa a MongoDB")
	return client
}

// GetCollection obtiene una colección específica de MongoDB
func GetCollection(client *mongo.Client, cfg *Config, collectionName string) *mongo.Collection {
	return client.Database(cfg.DBName).Collection(collectionName)
}
