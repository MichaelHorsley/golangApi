package participant

import (
	"gopkg.in/mgo.v2/bson"
)

type Participant struct {
	ID    bson.ObjectId `bson:"_id"`
	Name  string        `json:"name"`
	Email string        `json:"email"`
}

type Participants []Participant
