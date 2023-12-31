package router

import (
	"github.com/gorilla/mux"
	"github.com/juliofernandolepore/webserver/services"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/posts", services.GetAllPosts)
	router.HandleFunc("/posts/{id}", services.GetPost).Methods("GET")
	router.HandleFunc("/posts", services.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", services.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", services.DeletePost).Methods("DELETE")

	return router
}
