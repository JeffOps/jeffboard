package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aymerick/raymond"
)

var (
	homeTemplate *raymond.Template
)

func init() {
	homeTemplate, _ = raymond.ParseFile("views/index.html")
}

type Post struct {
	ID       uint      `db:"id"`
	Subject  string    `db:"subject"`
	Text     string    `db:"text"`
	Date     time.Time `db:"date_posted"`
	ThreadID uint      `db:"thread_id"`

	Drank uint `db:"drank"`
	Arank uint `db:"arank"`

	ThreadLastPosted time.Time `db:"last_post_in_thread"`
}

type Thread struct {
	ID    uint
	Posts []Post
}

type Homepage struct {
	Title   string
	Threads []Thread
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Queryx(`
	SELECT f.* FROM (
		SELECT subject, text, thread_id, date_posted,
		rank() OVER (PARTITION BY thread_id ORDER BY date_posted DESC) AS drank, -- last post in thread has drank = 1
		rank() OVER (PARTITION BY thread_id ORDER BY date_posted ASC) AS arank,  -- first post in thread has arank = 1
		max(date_posted) OVER (PARTITION BY thread_id) AS last_post_in_thread FROM post
	) AS f 
	WHERE drank < 3 -- last two posts in each thread
	OR arank = 1 -- first post
	ORDER BY last_post_in_thread DESC, thread_id, arank, drank DESC;
	`)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	homepage := Homepage{
		Title: "Frontpage",
	}
	var currentThread *Thread
	for rows.Next() {
		var post Post
		err = rows.StructScan(&post)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if currentThread != nil && currentThread.ID != post.ThreadID {
			homepage.Threads = append(homepage.Threads, *currentThread)
			currentThread = nil
		}
		if currentThread == nil {
			currentThread = &Thread{ID: post.ThreadID}
		}
		currentThread.Posts = append(currentThread.Posts, post)
	}
	homepage.Threads = append(homepage.Threads, *currentThread)

	v := homeTemplate.MustExec(homepage)
	fmt.Fprintf(w, v)
}
