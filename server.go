package main

import (
	"net/http"
	"todos_backend/common"
	"todos_backend/routers"

	"github.com/urfave/negroni"
)

func main() {

	common.InitKeys()

	r := routers.InitRoutes()

	n := negroni.Classic() // Includes some default middle-wares
	n.UseHandler(r)

	// Fire up the server
	http.ListenAndServe("localhost:4000", n)
}
