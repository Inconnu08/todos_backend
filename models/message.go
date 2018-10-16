package models

import "gopkg.in/mgo.v2/bson"

type (
	// Message represents one of the structure of our resources
	Message struct {
		Id      bson.ObjectId `json:"id" bson:"_id"`
		Message string        `json:"message" bson:"message"`
	}
)
