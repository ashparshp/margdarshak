package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	// application config
	var app application

	// read from cmd line

	// connect to the db

	app.Domain = "example.com"

	// log server start
	log.Printf("Starting server on port %d", port)

	// start api server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	if err != nil {
		log.Fatal(err)
	}
}
