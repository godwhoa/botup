package main

import (
	"database/sql"
	"github.com/godwhoa/random-shit/botup.me/api"
	"github.com/godwhoa/random-shit/botup.me/postgres"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

const (
	SRV_ADDR = ":8080"
	DB_ADDR  = "postgres://postgres:bingbong@localhost/botup?sslmode=disable"
	CSALT    = "$2a$06$Ke2pAMZuWgu2tloy5RkjCu0rwbDjYFwkR2wx8AzzNtPMY1BWsVfB6"
)

func main() {
	db, err := sql.Open("postgres", DB_ADDR)
	if err != nil {
		log.Fatal("postgres.open", err)
	}
	store := sessions.NewCookieStore([]byte(CSALT))

	userservice := postgres.UserService{db}
	user_api := api.User{userservice, store}

	r := mux.NewRouter()
	r.HandleFunc("/api/register", user_api.Register).Methods("POST")
	r.HandleFunc("/api/login", user_api.Login).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	http.Handle("/", r)

	log.Printf("Listening on %s\n", SRV_ADDR)
	log.Fatal("ListenAndServe: ", http.ListenAndServe(SRV_ADDR, nil))
}
