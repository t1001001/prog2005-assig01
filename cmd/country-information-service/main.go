package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	// setting the port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("Port is not set - Setting port to default: " + port)
	}

	//starting the server
	log.Println("Starting server on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
