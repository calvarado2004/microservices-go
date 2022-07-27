package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	gRpcPort = "50001"
)

var client *mongo.Client

type Config struct {
}

func main() {

	mongoClient, err := connectToMongo()

	if err != nil {
		log.Panic(err)
	}

}

func connectToMongo() (*mongo.Client, error) {

	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")

	clientOptions := options.Client().ApplyURI(mongoURL)

	clientOptions.SetAuth(options.Credential{
		Username: user,
		Password: password,
	})

	connect, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Println("Error connecting to mongo:", err)
		return nil, err
	}

	return connect, nil

}
