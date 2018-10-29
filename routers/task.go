package routers

import (
	"github.com/julienschmidt/httprouter"
	"todos_backend/common"
	"todos_backend/controllers"
)

func SetTaskRoutes(router *httprouter.Router) *httprouter.Router {

	// Get a TaskController instance
	tc := controllers.NewTaskController(common.GetSession())

	// End points
	router.GET("/task/:id", tc.GetTask)
	router.GET("/tasks", tc.GetTasks)
	router.POST("/task", tc.CreateTask)
	router.DELETE("/task/:id", tc.RemoveTask)

	return router
}
