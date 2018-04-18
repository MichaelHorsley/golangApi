package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	router := NewRouter()
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
