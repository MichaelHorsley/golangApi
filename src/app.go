package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"poolgolang/src/infrastructure"
	"poolgolang/src/participant"
	"poolgolang/src/league"
	"poolgolang/src/round"
	"poolgolang/src/match"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter().StrictSlash(true)

	var allRoutes infrastructure.Routes

	allRoutes = append(allRoutes, participant.GetRoutes()...)
	allRoutes = append(allRoutes, league.GetRoutes()...)
	allRoutes = append(allRoutes, round.GetRoutes()...)
	allRoutes = append(allRoutes, match.GetRoutes()...)

	for _, route := range allRoutes {
		var handler http.Handler
		handler = route.HandlerFunc
		a.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
}

func (a *App) Run(addr string) {
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(addr, handlers.CORS(allowedOrigins, allowedMethods)(a.Router)))
}