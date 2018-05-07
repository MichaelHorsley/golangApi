package participant

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct{}

// SERVER the DB server
const SERVER = "localhost:27017"

// DBNAME the name of the DB instance
const DBNAME = "poolLeagues"

// DOCNAME the name of the document
const DOCNAME = "participants"

// GetParticipants returns the list of Participants
func (r Repository) GetParticipants() Participants {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(DOCNAME)

	results := Participants{}

	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

// AddParticipant inserts an Participant in the DB
func (r Repository) AddParticipant(participant Participant) bool {
	session, err := mgo.Dial(SERVER)

	defer session.Close()

	participant.ID = bson.NewObjectId()

	session.DB(DBNAME).C(DOCNAME).Insert(participant)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

// UpdateParticipant updates an Participant in the DB (not used for now)
func (r Repository) UpdateParticipant(participant Participant) bool {
	session, err := mgo.Dial(SERVER)

	defer session.Close()

	session.DB(DBNAME).C(DOCNAME).UpdateId(participant.ID, participant)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

// DeleteParticipant deletes an Participant (not used for now)
func (r Repository) DeleteParticipant(id string) string {
	session, err := mgo.Dial(SERVER)

	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		return "NOT FOUND"
	}

	oid := bson.ObjectIdHex(id)

	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}

	return "OK"
}
