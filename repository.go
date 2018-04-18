package main

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
const DOCNAME = "leagues"

// GetLeagues returns the list of Leagues
func (r Repository) GetLeagues() Leagues {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	results := Leagues{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	return results
}

// AddLeague inserts an League in the DB
func (r Repository) AddLeague(league League) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	league.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(league)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// UpdateLeague updates an League in the DB (not used for now)
func (r Repository) UpdateLeague(league League) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	session.DB(DBNAME).C(DOCNAME).UpdateId(league.ID, league)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// DeleteLeague deletes an League (not used for now)
func (r Repository) DeleteLeague(id string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return "NOT FOUND"
	}
	// Grab id
	oid := bson.ObjectIdHex(id)
	// Remove user
	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}
	// Write status
	return "OK"
}
