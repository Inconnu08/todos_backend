package common

import (
	"gopkg.in/mgo.v2"
	"log"
)

// getSession creates a new Mongo session and panics if connection error occurs
func GetSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	s.SetMode(mgo.Monotonic, true)

	return s
}
