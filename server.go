package main

import (
	"github.com/urfave/negroni"
	"todos_backend/controllers"

	// Standard library packages
	"net/http"

	// Third party packages
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	// Get a TaskController instance
	tc := controllers.NewTaskController(getSession())

	// Get a user resource
	r.GET("/user/:id", uc.GetUser)

	// Get all user resources
	r.GET("/users", uc.GetUsers)

	// Create a new user
	r.POST("/user", uc.CreateUser)

	// Remove an existing user
	r.DELETE("/user/:id", uc.RemoveUser)

	// Get a task resource
	r.GET("/task/:id", tc.GetTask)

	// Create a task resource
	r.POST("/task", tc.CreateTask)

	// Remove an existing task
	r.DELETE("/task/:id", tc.RemoveTask)

	n := negroni.Classic() // Includes some default middle-wares
	n.UseHandler(r)

	// Fire up the server
	http.ListenAndServe("localhost:4000", n)
}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}