package routes

import (
	"api/handlers"

	"github.com/gorilla/mux"
)

// NewRouter is creating user on new path user
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	/* Users Routes */
	r.HandleFunc("/users", handlers.PostUser).Methods("POST")
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.PutUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	/* Post Routes */
	r.HandleFunc("/posts", handlers.PostPost).Methods("POST")
	r.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", handlers.GetPost).Methods("GET")
	r.HandleFunc("/posts/{id}", handlers.PutPost).Methods("PUT")
	r.HandleFunc("/posts/{id}", handlers.DeletePost).Methods("DELETE")

	/* Feedback Routes */
	r.HandleFunc("/feedbacks", handlers.PostFeedback).Methods("POST")
	r.HandleFunc("/feedbacks", handlers.GetFeedbacks).Methods("GET")
	r.HandleFunc("/feedbacks/{id}", handlers.GetFeedback).Methods("GET")
	return r
}
