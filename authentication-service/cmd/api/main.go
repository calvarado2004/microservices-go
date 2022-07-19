package main

import (
	"calvarado2004/microservices-go/authentication/data"
	"database/sql"
	"log"
)

const webPort = "80"

type Config struct {
	DB *sql.DB
	Models data.Models
}


func main {
	log.Printn("Starting authentication service on port", webPort)


	app := Config{}

	srv := &http.Server{

		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
	
}