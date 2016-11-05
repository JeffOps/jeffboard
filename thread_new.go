package main

import (
	"fmt"
	"net/http"

	"github.com/aymerick/raymond"
)

var (
	newThreadTemplate *raymond.Template
)

func init() {
	newThreadTemplate, _ = raymond.ParseFile("views/new-thread.html")
	registerLayout(newThreadTemplate)
}

func handleNewThread(w http.ResponseWriter, r *http.Request) {
	v := newThreadTemplate.MustExec(map[string]interface{}{
		"page":  "new-thread",
		"title": "jeffboard - new thread",
	})
	fmt.Fprintf(w, v)
}

func handleNewThreadPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
