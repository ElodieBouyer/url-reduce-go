package main

import (
	"log"
	"net/http"

	"app/handlers"
)

func main() {
	// Add handlers.
	http.HandleFunc("/add", handlers.Add)
	http.HandleFunc("/get", handlers.Get)

	// Start the server.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
