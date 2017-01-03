package api

import (
	"net/http"
	"strconv"

	"github.com/godwhoa/random-shit/botup.me/botup"
	"github.com/gorilla/sessions"
)

// User form validation
func validate_reg_form(r *http.Request) (botup.User, error) {
	email := r.FormValue("email")
	user := r.FormValue("user")
	pass := r.FormValue("pass")
	if email == "" || user == "" || pass == "" {
		return botup.User{}, invalidForm
	}
	return botup.User{Email: email, User: user, Pass: pass}, nil
}

// Bot form validation
func validate_addbot_form(r *http.Request, store *sessions.CookieStore) (botup.Bot, error) {
	session, _ := store.Get(r, "login")

	addr := r.FormValue("addr")
	channel := r.FormValue("channel")
	if addr == "" || channel == "" {
		return botup.Bot{}, invalidForm
	}

	bot := botup.Bot{}
	bot.UID = session.Values["uid"].(string)
	bot.Addr = addr
	bot.Channel = channel

	return bot, nil
}

func validate_addplugin_form(r *http.Request, store *sessions.CookieStore) (botup.Plugin, error) {
	session, _ := store.Get(r, "login")

	plugin := botup.Plugin{}

	bid, err := strconv.Atoi(r.FormValue("bid"))
	plugin_ := r.FormValue("plugin")

	if plugin_ == "" || err != nil {
		return plugin, invalidForm
	}

	plugin.BID = bid
	plugin.UID = session.Values["uid"].(string)
	plugin.Plugin = plugin_

	return plugin, nil
}

func validate_removebot_form(r *http.Request, store *sessions.CookieStore) (botup.Bot, error) {
	session, _ := store.Get(r, "login")

	bid, err := strconv.Atoi(r.FormValue("bid"))
	if err != nil {
		return botup.Bot{}, invalidForm
	}

	bot := botup.Bot{}
	bot.UID = session.Values["uid"].(string)
	bot.BID = bid

	return bot, nil
}

func validate_removeplugin_form(r *http.Request, store *sessions.CookieStore) (botup.Plugin, error) {
	session, _ := store.Get(r, "login")

	bid, err := strconv.Atoi(r.FormValue("bid"))
	plugin := r.FormValue("plugin")
	if plugin == "" || err != nil {
		return botup.Plugin{}, invalidForm
	}

	return botup.Plugin{bid, session.Values["uid"].(string), plugin}, nil
}
