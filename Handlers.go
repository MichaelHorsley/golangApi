package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func LeagueIndex(w http.ResponseWriter, r *http.Request) {
	leagues := Leagues{
		League{Name: "Write presentation"},
		League{Name: "Host meetup"},
	}

	if err := json.NewEncoder(w).Encode(leagues); err != nil {
		panic(err)
	}
}

func LeagueShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	leagueId := vars["leagueId"]
	fmt.Fprintln(w, "League show:", leagueId)
}
