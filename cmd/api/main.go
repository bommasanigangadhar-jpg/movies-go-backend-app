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

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World Hello")
}

func main() {
	// set application config
	var app application
	// read from command line
	//connect to the database
	app.Domain = "example.com"

	log.Println("Starting application on port", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	if err != nil {
		log.Fatal(err)
	}

}
