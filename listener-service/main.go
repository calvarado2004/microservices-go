package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// try to connect to rabbitmq
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()
	log.Println("Connected to RabbitMQ")

	// start listening for messages

	// create a consumer

	// watch for messages on queue and consume events

}

func connect() (*amqp.Connection, error) {

	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// don't continue until rabbitmq is up

	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			fmt.Println("RabbitMQ is not up yet.")
			counts++
		} else {
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("Waiting for RabbitMQ to come up. Back off for", backOff)
		time.Sleep(backOff)

	}

	return connection, nil

}
