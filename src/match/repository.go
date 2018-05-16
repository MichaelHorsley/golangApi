package match

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
const DOCNAME = "matches"

// GetMatches returns the list of Matches
func (r Repository) GetMatches() Matches {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(DOCNAME)

	results := Matches{}

	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

// AddMatch inserts an Match in the DB
func (r Repository) AddMatch(match Match) bool {
	session, err := mgo.Dial(SERVER)

	defer session.Close()

	match.ID = bson.NewObjectId()

	session.DB(DBNAME).C(DOCNAME).Insert(match)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

// UpdateMatch updates an Match in the DB (not used for now)
func (r Repository) UpdateMatch(match Match) bool {
	session, err := mgo.Dial(SERVER)

	defer session.Close()

	session.DB(DBNAME).C(DOCNAME).UpdateId(match.ID, match)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

// DeleteMatch deletes an Match (not used for now)
func (r Repository) DeleteMatch(id string) string {
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
