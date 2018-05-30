package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func (a *App) Run(addr string) {
	a.Router = NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(allowedOrigins, allowedMethods)(a.Router)))
}
