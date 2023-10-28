package router

import (
	"github.com/gorilla/mux"
	"github.com/juliofernandolepore/webserver/services"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/posts", services.GetAllPosts)
	return router
}
