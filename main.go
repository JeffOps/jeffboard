package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ConnectedVentures/gonfigurator"
	"github.com/aymerick/raymond"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgreSQLConfig struct {
	User     string
	Host     string
	Password string
	Database string
	Port     uint
}

type Config struct {
	PostgreSQL PostgreSQLConfig
}

var (
	config Config
	db     *sqlx.DB
)

func registerLayout(tpl *raymond.Template) {
	tpl.RegisterPartialFiles("views/header.html", "views/footer.html")
}

func initDb(pgConfig PostgreSQLConfig) {
	if pgConfig.Port == 0 {
		pgConfig.Port = 5432
	}
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d", pgConfig.User, pgConfig.Password, pgConfig.Database, pgConfig.Host, pgConfig.Port)
	var err error
	db, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("Cannot connect to database: %s", err.Error())
	}
}

var (
	headerTemplate, footerTemplate *raymond.Template
)

func init() {
	headerTemplate, _ = raymond.ParseFile("views/header.html")
	footerTemplate, _ = raymond.ParseFile("views/header.html")
}

func main() {
	gonfigurator.Parse("config.yml", &config)
	initDb(config.PostgreSQL)
	router := mux.NewRouter().StrictSlash(true)
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
