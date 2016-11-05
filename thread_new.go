package main

import (
	"fmt"
	"net/http"
)

func handleNewThread(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func handleNewThreadPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
