package server

import (
	"log"
	"net/http"

	userHandler "github.com/M1n007/go-rest-api/repository/users"
	"github.com/gorilla/mux"
)

// Init Server
func Init() {
	router := mux.NewRouter()

	//Routes
	router.HandleFunc("/api/users/v1", userHandler.ReturnAllUsers).Methods("GET")
	router.HandleFunc("/api/users/v1", userHandler.CreateUser).Methods("POST")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":9000", router))
}
