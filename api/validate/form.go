package validate

import (
	"errors"

	"github.com/godwhoa/random-shit/botup.me/botup"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
	"strconv"
)

const (
	NoEmpty = iota
	IsInt
	IsString
)

type Rules []int

var ValidationErr = errors.New("Form field failed one or more rules")

func Validate(r *http.Request, isParam bool, formRules map[string]Rules) (map[string]interface{}, error) {
	parsed := make(map[string]interface{})
	vars := mux.Vars(r)

	for field, rules := range formRules {
		var value string
		value = r.FormValue(field)
		if isParam {
			value = vars[field]
		}
		for _, rule := range rules {
			switch rule {
			case NoEmpty, IsString:
				if value == "" {
					return parsed, ValidationErr
				}
				parsed[field] = value
			case IsInt:
				i, err := strconv.Atoi(value)
				if err != nil {
					return parsed, ValidationErr
				}
				parsed[field] = int(i)
			}
		}
	}
	return parsed, nil
}

// Form validation rules
var (
	RegRules = map[string]Rules{
		"email": {IsString, NoEmpty},
		"user":  {IsString, NoEmpty},
		"pass":  {IsString, NoEmpty},
	}
	LoginRules = map[string]Rules{
		"email": {IsString, NoEmpty},
		"pass":  {IsString, NoEmpty},
	}
	AddBotRules = map[string]Rules{
		"nick":    {IsString, NoEmpty},
		"addr":    {IsString, NoEmpty},
		"channel": {IsString, NoEmpty},
	}
	AddPluginRules = map[string]Rules{
		"bid":    {IsInt},
		"plugin": {IsString, NoEmpty},
	}
	GetBotRules = map[string]Rules{
		"BID": {IsInt},
	}
	RemoveBotRules = map[string]Rules{
		"addr":    {IsString, NoEmpty},
		"channel": {IsString, NoEmpty},
	}
	RemovePluginRules = map[string]Rules{
		"bid":    {IsInt, NoEmpty},
		"plugin": {IsString, NoEmpty},
	}
)

func Registration(r *http.Request) (botup.User, error) {
	form, err := Validate(r, false, RegRules)
	if err != nil {
		return botup.User{}, ValidationErr
	}
	return botup.User{
		Email: form["email"].(string),
		User:  form["user"].(string),
		Pass:  form["pass"].(string),
	}, nil
}

func Login(r *http.Request) (botup.User, error) {
	form, err := Validate(r, false, LoginRules)
	if err != nil {
		return botup.User{}, ValidationErr
	}
	return botup.User{
		Email: form["email"].(string),
		Pass:  form["pass"].(string),
	}, nil
}

func AddBot(r *http.Request, store *sessions.CookieStore) (botup.Bot, error) {
	session, _ := store.Get(r, "login")

	form, err := Validate(r, false, AddBotRules)
	if err != nil {
		return botup.Bot{}, ValidationErr
	}

	return botup.Bot{
		UID:     session.Values["uid"].(string),
		Addr:    form["addr"].(string),
		Nick:    form["nick"].(string),
		Channel: form["channel"].(string),
		Alive:   true,
	}, nil
}

func AddPlugin(r *http.Request, store *sessions.CookieStore) (botup.Plugin, error) {
	session, _ := store.Get(r, "login")

	form, err := Validate(r, false, AddPluginRules)
	if err != nil {
		return botup.Plugin{}, ValidationErr
	}

	return botup.Plugin{
		UID:    session.Values["uid"].(string),
		BID:    form["bid"].(int),
		Plugin: form["plugin"].(string),
	}, nil
}

func GetBots(r *http.Request, store *sessions.CookieStore) (string, error) {
	session, _ := store.Get(r, "login")
	return session.Values["uid"].(string), nil
}

func GetBot(r *http.Request, store *sessions.CookieStore) (string, int, error) {
	session, _ := store.Get(r, "login")

	form, err := Validate(r, true, GetBotRules)
	if err != nil {
		return "", -1, ValidationErr
	}
	return session.Values["uid"].(string), form["BID"].(int), nil
}

func GetAllPlugins(r *http.Request, store *sessions.CookieStore) (string, error) {
	session, _ := store.Get(r, "login")
	return session.Values["uid"].(string), nil
}

func GetPlugins(r *http.Request, store *sessions.CookieStore) (string, int, error) {
	session, _ := store.Get(r, "login")

	form, err := Validate(r, true, GetBotRules)
	if err != nil {
		return "", -1, ValidationErr
	}
	return session.Values["uid"].(string), form["BID"].(int), nil
}

func RemoveBot(r *http.Request, store *sessions.CookieStore) (botup.Bot, error) {
	session, _ := store.Get(r, "login")

	form, err := Validate(r, false, RemoveBotRules)
	if err != nil {
		return botup.Bot{}, ValidationErr
	}

	return botup.Bot{
		UID: session.Values["uid"].(string),
		BID: form["bid"].(int),
	}, nil
}

func RemovePlugin(r *http.Request, store *sessions.CookieStore) (botup.Plugin, error) {
	session, _ := store.Get(r, "login")

	form, err := Validate(r, false, RemovePluginRules)
	if err != nil {
		return botup.Plugin{}, ValidationErr
	}

	return botup.Plugin{
		UID:    session.Values["uid"].(string),
		BID:    form["bid"].(int),
		Plugin: form["plugin"].(string),
	}, nil
}
