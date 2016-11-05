package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/aymerick/raymond"
	"github.com/gorilla/mux"
)

var (
	threadTemplate *raymond.Template
)

func init() {
	threadTemplate, _ = raymond.ParseFile("views/thread.html")
}

func handleThread(w http.ResponseWriter, r *http.Request) {
	var posts []Post

	strID := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(strID)

	err := db.Select(&posts, "SELECT * FROM post WHERE thread_id = $1", id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	v := threadTemplate.MustExec(map[string]interface{}{
		"id":    posts[0].ThreadID,
		"posts": posts,
	})
	fmt.Fprintf(w, v)
}

func handleThreadPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
