package round

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

//Controller ...
type Controller struct {
	Repository Repository
}

// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	rounds := c.Repository.GetRounds()

	log.Println(rounds)

	data, _ := json.Marshal(rounds)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)

	w.Write(data)

	return
}

// AddRound POST /
func (c *Controller) AddRound(w http.ResponseWriter, r *http.Request) {
	var round Round

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request

	if err != nil {
		log.Fatalln("Error AddRound", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddRound", err)
	}

	if err := json.Unmarshal(body, &round); err != nil { // unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddRound unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	success := c.Repository.AddRound(round)

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	return
}

// UpdateRound PUT /
func (c *Controller) UpdateRound(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var round Round

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request

	if err != nil {
		log.Fatalln("Error UpdateRound", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddaUpdateRoundlbum", err)
	}

	if err := json.Unmarshal(body, &round); err != nil { // unmarshall body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error UpdateRound unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	round.ID = bson.ObjectIdHex(id)

	success := c.Repository.UpdateRound(round) // updates the Round in the DB

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)

	return
}

// DeleteRound DELETE /
func (c *Controller) DeleteRound(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.Repository.DeleteRound(id); err != "" { // delete a Round by id
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	return
}
