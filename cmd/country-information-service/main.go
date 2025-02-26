package main

import (
	"log"
	"net/http"
	"os"

	"github.com/t1001001/prog2005-assignment-01/internal/constants"
	"github.com/t1001001/prog2005-assignment-01/internal/handlers"
)

func main() {
	// setting the port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("Port is not set - Setting port to default: " + port)
	}

	// setting up the endpoints
	http.HandleFunc(constants.DEFAULT_PATH, handlers.DefaultHandler)
	http.HandleFunc(constants.INFO_PATH+"/", handlers.InfoHandler)
	http.HandleFunc(constants.POPULATION_PATH+"/", handlers.PopulationHandler)
	http.HandleFunc(constants.STATUS_PATH, handlers.StatusHandler)

	// starting server
	log.Println("Starting server on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
