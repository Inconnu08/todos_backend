package common

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"taskmanager/Godeps/_workspace/src/github.com/dgrijalva/jwt-go"
)


// Middleware (this function) makes adding more than one layer of middleware easy
// by specifying them as a list. It will run the last specified handler first.
func Middleware(h httprouter.Handle, middleware ...func(httprouter.Handle) httprouter.Handle) httprouter.Handle {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}

// Middleware for authorizing only with valid jwt token.
func AuthMiddleware(next httprouter.Handle) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// validate the token
		token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
		if err == nil && token.Valid {
			// allow to handler
			next(w, r, ps)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Authentication failed")
		}
	}
}
