package main

import (
	"calvarado2004/microservices-go/log-service/data"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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
	Models data.Models
}

func main() {

	mongoClient, err := connectToMongo()

	if err != nil {
		log.Panic(err)
	}

	client = mongoClient

	//create a mongo context in order to be able to disconnect

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//close connection on exit

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	// start the server
	//go app.serve()
	log.Println("Starting server on port", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

func (app *Config) serve() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
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

	log.Println("Connected to mongo")

	return connect, nil

}
