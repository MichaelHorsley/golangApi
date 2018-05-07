package league

import (
	"gopkg.in/mgo.v2/bson"
)

type League struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `json:"name"`
}

type Leagues []League
