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
	login_cache := make(map[string]string)

	userservice, botservice := postgres.UserService{db}, postgres.BotService{db}
	user_api := api.User{userservice, store, login_cache}
	bot_api := api.Bot{botservice, store}

	r := mux.NewRouter()
	r.HandleFunc("/api/user/register", user_api.Register).Methods("POST")
	r.HandleFunc("/api/user/login", user_api.Login).Methods("POST")
	r.HandleFunc("/api/user/logout", user_api.Logout).Methods("GET")

	r.HandleFunc("/api/bot/add", bot_api.AddBot).Methods("POST")
	r.HandleFunc("/api/bot/remove", bot_api.RemoveBot).Methods("POST")

	r.HandleFunc("/api/plugin/add", bot_api.AddPlugin).Methods("POST")
	r.HandleFunc("/api/plugin/remove", bot_api.RemovePlugin).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	http.Handle("/", r)

	log.Printf("Listening on %s\n", SRV_ADDR)
	log.Fatal("ListenAndServe: ", http.ListenAndServe(SRV_ADDR, nil))
}
