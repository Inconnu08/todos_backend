package controllers

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"taskmanager/common"
	"todos_backend/models"

	"github.com/julienschmidt/httprouter"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// UserController represents the controller for operating on the User resource
	UserController struct {
		session *mgo.Session
	}
)

// NewUserController provides a reference to a UserController with provided mongo session
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// GetUser retrieves an individual user resource
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub user
	u := models.User{}

	// Fetch user
	if err := uc.session.DB("todos").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// GetUser retrieves all user resources
func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var ul []models.User

	// Fetch user
	if err := uc.session.DB("todos").C("users").Find(nil).All(&ul); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(ul)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// CreateUser creates a new user resource
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an user to be populated from the body
	u := models.User{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&u)

	// Add an Id
	u.Id = bson.NewObjectId()

	hPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	u.HashPassword = hPass
	// clear the incoming text password
	u.Password = ""

	// Write the user to mongo
	err = uc.session.DB("todos").C("users").Insert(&u)

	// clear hashed password
	u.HashPassword = nil

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var dataResource LoginModel
	var token string
	var u models.User

	//fmt.Println("Login handler")
	//body, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(string(body))
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		fmt.Println("Decoding error...")
		w.WriteHeader(500)
		return
	}

	fmt.Println(dataResource)
	loginUser := models.User{
		Email:    dataResource.Email,
		Password: dataResource.Password,
	}
	fmt.Printf("%s right?", loginUser.Email)
	err = uc.session.DB("todos").C("users").Find(bson.M{"email": loginUser.Email}).One(&u)
	fmt.Println(u)
	if err != nil {
		fmt.Println("Error finding user...")
		w.WriteHeader(500)
		return
	}
	// Validate password
	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(dataResource.Password))
	fmt.Printf("\n%s\n %s\n", u.HashPassword, loginUser.Password)
	if err == nil {
		//if login is successful

		// Generate JWT token
		fmt.Printf("Generate token...\n")
		fmt.Println(u.Email)
		token, err = common.GenerateJWT(u.Email, "member")
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		// Clean-up the hashpassword to eliminate it from response JSON
		u.HashPassword = nil
		authUser := AuthUserModel{
			User:  u,
			Token: token,
		}
		j, err := json.Marshal(AuthUserResource{Data: authUser})
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)

	} else {
		w.WriteHeader(401)
		return
	}
}

// RemoveUser removes an existing user resource
func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err := uc.session.DB("todos").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// Write status
	w.WriteHeader(200)
}
