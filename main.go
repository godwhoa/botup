package main

import (
	"database/sql"
	"github.com/godwhoa/random-shit/botup.me/api"
	"github.com/godwhoa/random-shit/botup.me/postgres"
	"github.com/godwhoa/random-shit/botup.me/utils/decorators"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"github.com/uber-go/zap"
	"net/http"
)

const (
	SRV_ADDR = ":8080"
	DB_ADDR  = "user=postgres dbname=botup password=bingbong sslmode=disable"
	CSALT    = "$2a$06$Ke2pAMZuWgu2tloy5RkjCu0rwbDjYFwkR2wx8AzzNtPMY1BWsVfB6"
)

func main() {
	// Init logger
	log := zap.New(
		zap.NewJSONEncoder(
			zap.RFC3339Formatter("@timestamp"),
			zap.MessageKey("@message"),
			zap.LevelString("@level"),
		),
	)

	// Init postgres connection
	db, err := sql.Open("postgres", DB_ADDR)
	if err != nil {
		log.Fatal("postgres.open", zap.Error(err))
	}

	// Init cookie and cache stores
	store := sessions.NewCookieStore([]byte(CSALT))
	login_cache := make(map[string]string)

	// Init APIs
	userservice, botservice := postgres.UserService{db}, postgres.BotService{db}
	user_api := api.User{userservice, store, login_cache, log}
	bot_api := api.Bot{botservice, store, log}

	r := mux.NewRouter()
	// User handlers
	r.HandleFunc("/api/user/register", user_api.Register).Methods("POST")
	r.HandleFunc("/api/user/login", user_api.Login).Methods("POST")
	r.HandleFunc("/api/user/logout", user_api.Logout).Methods("GET")

	// Bot handlers
	r.HandleFunc("/api/bot/add",
		decorators.Auth(bot_api.AddBot, store, login_cache)).Methods("POST")
	r.HandleFunc("/api/bot/get/{BID}",
		decorators.Auth(bot_api.GetBot, store, login_cache)).Methods("GET")
	r.HandleFunc("/api/bot/getall",
		decorators.Auth(bot_api.GetBots, store, login_cache)).Methods("GET")

	// Plugin handlers
	r.HandleFunc("/api/plugin/add",
		decorators.Auth(bot_api.AddPlugin, store, login_cache)).Methods("POST")
	r.HandleFunc("/api/plugin/get/{BID}",
		decorators.Auth(bot_api.GetPlugins, store, login_cache)).Methods("GET")
	r.HandleFunc("/api/plugin/getall",
		decorators.Auth(bot_api.GetAllPlugins, store, login_cache)).Methods("GET")

	// Public
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./client/")))
	http.Handle("/", r)

	log.Info("http.ListenAndServe", zap.String("addr", SRV_ADDR))
	log.Fatal("http.ListenAndServe: ", zap.Error(http.ListenAndServe(SRV_ADDR, nil)))
}
