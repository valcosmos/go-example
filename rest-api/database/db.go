package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func GetCollection(name string) *mongo.Collection {
	return MongoClient.Database("test").Collection(name)
}

func StartMongoDB() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set you 'MONGODB_URI'")
	}
	var err error
	MongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connect to mongodb")
	return nil
}

func CloseMongoDB() {
	err := MongoClient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
}
