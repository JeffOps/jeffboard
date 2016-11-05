package main

import (
	"fmt"
	"net/http"
)

func handleThread(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func handleThreadPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
