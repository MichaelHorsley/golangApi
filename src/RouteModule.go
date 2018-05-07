package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"poolgolang/src/infrastructure"
	"poolgolang/src/league"
	"poolgolang/src/participant"
)

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
