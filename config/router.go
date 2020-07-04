package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//CreateRouter CreateRouter
func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

//RunServer RunServer
func RunServer(router *mux.Router) {
	port := "localhost:8000"
	fmt.Println("Setting Web Server at port : " + port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
