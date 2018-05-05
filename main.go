package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"poolgolang/infrastructure"
	"poolgolang/league"
	"poolgolang/participant"
)

func main() {
	var foo = NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(allowedOrigins, allowedMethods)(foo)))
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	var allRoutes infrastructure.Routes

	allRoutes = append(allRoutes, participant.GetRoutes()...)
	allRoutes = append(allRoutes, league.GetRoutes()...)

	for _, route := range allRoutes {
		var handler http.Handler
		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
