package routers

import "github.com/julienschmidt/httprouter"

func InitRoutes() *httprouter.Router {

	// Initialize router
	router := httprouter.New()
	// Routes for the User entity
	router = SetUserRoutes(router)
	// Routes for the Task entity
	router = SetTaskRoutes(router)

	return router
}
