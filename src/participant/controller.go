package participant

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
	participants := c.Repository.GetParticipants()

	log.Println(participants)

	data, _ := json.Marshal(participants)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)

	w.Write(data)

	return
}

// AddParticipant POST /
func (c *Controller) AddParticipant(w http.ResponseWriter, r *http.Request) {
	var participant Participant

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request

	if err != nil {
		log.Fatalln("Error AddParticipant", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddParticipant", err)
	}

	if err := json.Unmarshal(body, &participant); err != nil {
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddParticipant unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	success := c.Repository.AddParticipant(participant)

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	return
}

// UpdateParticipant PUT /
func (c *Controller) UpdateParticipant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var participant Participant

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request

	if err != nil {
		log.Fatalln("Error UpdateParticipant", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddaUpdateParticipantlbum", err)
	}

	if err := json.Unmarshal(body, &participant); err != nil { // unmarshall body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error UpdateParticipant unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	participant.ID = bson.ObjectIdHex(id)

	success := c.Repository.UpdateParticipant(participant) // updates the Participant in the DB

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)

	return
}

// DeleteParticipant DELETE /
func (c *Controller) DeleteParticipant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.Repository.DeleteParticipant(id); err != "" { // delete a Participant by id
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
