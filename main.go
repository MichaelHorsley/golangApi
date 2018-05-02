package main

import (
	"log"
	"net/http"
	"poolgolang/participant"

	"github.com/gorilla/handlers"
)

func main() {
	router := participant.NewRouter()
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
