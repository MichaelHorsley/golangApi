package match

import (
	"gopkg.in/mgo.v2/bson"
)

type Match struct {
	ID    bson.ObjectId `bson:"_id"`
}

type Matches []Match
