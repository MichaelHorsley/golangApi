package round

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
const DOCNAME = "rounds"

// GetRounds returns the list of Rounds
func (r Repository) GetRounds() Rounds {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(DOCNAME)

	results := Rounds{}

	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

// AddRound inserts an Round in the DB
func (r Repository) AddRound(round Round) bool {
	session, err := mgo.Dial(SERVER)

	defer session.Close()

	round.ID = bson.NewObjectId()

	session.DB(DBNAME).C(DOCNAME).Insert(round)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

// UpdateRound updates an Round in the DB (not used for now)
func (r Repository) UpdateRound(round Round) bool {
	session, err := mgo.Dial(SERVER)

	defer session.Close()

	session.DB(DBNAME).C(DOCNAME).UpdateId(round.ID, round)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

// DeleteRound deletes an Round (not used for now)
func (r Repository) DeleteRound(id string) string {
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
