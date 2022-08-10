package main

import (
	"calvarado2004/microservices-go/listener-service/event"
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

	// start listening for messages
	log.Println("Listening and consuming RabbitMQ messages")

	// create a consumer
	consumer, err := event.NewConsumer(rabbitConn)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// watch for messages on queue and consume events
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR", "log.CRITICAL"})
	if err != nil {
		log.Println(err)
	}
}

func connect() (*amqp.Connection, error) {

	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// don't continue until rabbitmq is up

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err != nil {
			fmt.Println("RabbitMQ is not up yet.")
			counts++
		} else {
			log.Println("Connected to RabbitMQ")
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
