package routes

import (
	"api/handlers"

	"github.com/gorilla/mux"
)

// NewRouter is creating user on new path user
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/users", handlers.PostUser).Methods("POST")
	return r
}
