package routers

import (
	"github.com/julienschmidt/httprouter"
	"todos_backend/common"
	"todos_backend/controllers"
)

func SetUserRoutes(router *httprouter.Router) *httprouter.Router {

	// Get a UserController instance
	uc := controllers.NewUserController(common.GetSession())

	router.GET("/user/:id", uc.GetUser)
	router.GET("/users", uc.GetUsers)
	router.POST("/register", uc.CreateUser)
	router.POST("/login", uc.Login)
	router.DELETE("/user/:id", uc.RemoveUser)

	return router
}
