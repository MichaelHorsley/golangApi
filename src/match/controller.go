package match

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
	matches := c.Repository.GetMatches()

	log.Println(matches)

	data, _ := json.Marshal(matches)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)

	w.Write(data)

	return
}

// AddMatch POST /
func (c *Controller) AddMatch(w http.ResponseWriter, r *http.Request) {
	var match Match

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request

	if err != nil {
		log.Fatalln("Error AddMatch", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddMatch", err)
	}

	if err := json.Unmarshal(body, &match); err != nil { // unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddMatch unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	success := c.Repository.AddMatch(match)

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	return
}

// UpdateMatch PUT /
func (c *Controller) UpdateMatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var match Match

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request

	if err != nil {
		log.Fatalln("Error UpdateMatch", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddaUpdateMatchlbum", err)
	}

	if err := json.Unmarshal(body, &match); err != nil { // unmarshall body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error UpdateMatch unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	match.ID = bson.ObjectIdHex(id)

	success := c.Repository.UpdateMatch(match) // updates the Match in the DB

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)

	return
}

// DeleteMatch DELETE /
func (c *Controller) DeleteMatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.Repository.DeleteMatch(id); err != "" { // delete a Match by id
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
