package models

import "gopkg.in/mgo.v2/bson"

type (
	// Task represents one of the structure of our resources
	Task struct {
		Id      bson.ObjectId `json:"id" bson:"_id"`
		Title   string        `json:"title" bson:"title"`
		Done    bool          `json:"done" bson:"done"`
		TaskFor User          `json:"task_for" bson:"task_for"`
	}
)
