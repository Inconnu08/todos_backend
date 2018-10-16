package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todos_backend/models"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// TaskController represents the controller for operating on the Task resource
	TaskController struct {
		session *mgo.Session
	}
)

// NewTaskController provides a reference to a TaskController with provided mongo session
func NewTaskController(s *mgo.Session) *TaskController {
	return &TaskController{s}
}

// GetTask retrieves an individual Task resource
func (uc TaskController) GetTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub Task
	t := models.Task{}

	// Fetch Task
	if err := uc.session.DB("todos").C("tasks").FindId(oid).One(&t); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	tj, _ := json.Marshal(t)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", tj)
}

// CreateTask creates a new Task resource
func (uc TaskController) CreateTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an Task to be populated from the body
	t := models.Task{}

	// Populate the Task data
	json.NewDecoder(r.Body).Decode(&t)

	// Add an Id
	t.Id = bson.NewObjectId()

	// Write the Task to mongo
	uc.session.DB("todos").C("tasks").Insert(t)

	// Marshal provided interface into JSON structure
	tj, _ := json.Marshal(t)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", tj)
}

// RemoveTask removes an existing Task resource
func (uc TaskController) RemoveTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove Task
	if err := uc.session.DB("todos").C("tasks").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// Write status
	w.WriteHeader(200)
}
