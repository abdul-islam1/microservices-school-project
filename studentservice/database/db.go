package database

import (
	"context"
	"log"
	"os"
	"time"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Database *mongo.Database

func Connect() error {
	// From Environment variables 
	// connectionString := os.Getenv("MONGODB_URI")
	connectionString := fmt.Sprintf(
	"mongodb://%s:%s@%s?authSource=admin",
	os.Getenv("DATABASE_USERNAME"),
	os.Getenv("DATABASE_PASSWORD"),
	os.Getenv("DATABASE_URL"),
)
	
	// if connectionString == "" {
	// 	connectionString = "mongodb://mongo:27017" // default value
	// }

	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		databaseName = "kindergarten" // default value
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return err
	}

	Client = client
	Database = client.Database("kindergarten")
	log.Println("Connected to MongoDB successfully!")
	return nil
}

func GetCollection(collectionName string) *mongo.Collection {
	return Database.Collection(collectionName)
}