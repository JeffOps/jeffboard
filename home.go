package main

import (
	"fmt"
	"net/http"

	"github.com/aymerick/raymond"
)

var (
	homeTemplate *raymond.Template
)

func init() {
	homeTemplate, _ = raymond.ParseFile("views/index.html")
}

type Post struct {
	ID         uint
	Text       string
	Attachment string
}

type Thread struct {
	Posts []Post
}

type Homepage struct {
	Title   string
	Threads []Thread
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	homepage := Homepage{
		Title: "Homepage",
		Threads: []Thread{
			Thread{
				Posts: []Post{
					Post{
						ID:         1,
						Text:       "Thread 1 - post 1",
						Attachment: "http://placekitten.com/200/200",
					},
					Post{
						ID:         2,
						Text:       "Thread 1 - post 2 (reply)",
						Attachment: "http://placekitten.com/200/200",
					},
				},
			},
			Thread{
				Posts: []Post{
					Post{
						ID:         3,
						Text:       "Thread 2 - post 1",
						Attachment: "http://placekitten.com/200/200",
					},
					Post{
						ID:         4,
						Text:       "Thread 2 - post 2 (reply)",
						Attachment: "http://placekitten.com/200/200",
					},
				},
			},
		},
	}
	v := homeTemplate.MustExec(homepage)
	fmt.Fprintf(w, v)
}
