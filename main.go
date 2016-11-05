package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleHome).Methods("GET")
	router.HandleFunc("/thread", handleNewThread).Methods("GET")
	router.HandleFunc("/thread", handleNewThreadPost).Methods("POST")
	router.HandleFunc("/thread/{id:[0-9]+}", handleThread).Methods("GET")
	// Post reply
	// This will be done directly from the thread
	router.HandleFunc("/thread/{id:[0-9]+}", handleThreadPost).Methods("POST")

	s := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	router.PathPrefix("/assets/").Handler(s)
	log.Fatal(http.ListenAndServe(":8900", router))
}
