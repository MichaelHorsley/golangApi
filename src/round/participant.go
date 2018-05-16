package round

import (
	"gopkg.in/mgo.v2/bson"
)

type Round struct {
	ID    bson.ObjectId `bson:"_id"`
}

type Rounds []Round
