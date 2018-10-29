package models

import "gopkg.in/mgo.v2/bson"

type (
	// User represents one of the structure of our resources
	User struct {
		Id           bson.ObjectId `json:"id" bson:"_id"`
		Name         string        `json:"name" bson:"name"`
		Email        string        `json:"email"`
		Gender       string        `json:"gender" bson:"gender"`
		Age          int           `json:"age" bson:"age"`
		Password     string        `json:"password,omitempty" bson:"password"`
		HashPassword []byte        `json:"hash_password,omitempty"  bson:"hash_password"`
		Votes        int           `json:"votes" bson:"votes"`
	}
)
