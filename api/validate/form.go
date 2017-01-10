package validate

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/godwhoa/random-shit/botup.me/botup"
	"github.com/gorilla/sessions"
)

const (
	NoEmpty = iota
	IsInt
	IsString
)

type Rules []int

var ValidationErr = errors.New("Form field failed one or more rules")

func Validate(r *http.Request, formRules map[string]Rules) (map[string]interface{}, error) {
	parsed := make(map[string]interface{})
	pass := false
formLoop:
	for field, rules := range formRules {
		value := r.FormValue(field)
		for _, rule := range rules {
			switch rule {
			case NoEmpty, IsString:
				if value == "" {
					break formLoop
				}
				parsed[field] = value
				pass = true
			case IsInt:
				i, err := strconv.Atoi(value)
				if err != nil {
					break formLoop
				}
				parsed[field] = i
				pass = true
			}
		}
	}
	if !pass {
		return parsed, ValidationErr
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
		"addr":    {IsString, NoEmpty},
		"channel": {IsString, NoEmpty},
	}
	AddPluginRules = map[string]Rules{
		"bid":    {IsInt, NoEmpty},
		"plugin": {IsString, NoEmpty},
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
	form, err := Validate(r, RegRules)
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
	form, err := Validate(r, LoginRules)
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

	form, err := Validate(r, AddBotRules)
	if err != nil {
		return botup.Bot{}, ValidationErr
	}

	return botup.Bot{
		UID:     session.Values["uid"].(string),
		Addr:    form["addr"].(string),
		Channel: form["channel"].(string),
		Alive:   true,
	}, nil
}

func AddPlugin(r *http.Request, store *sessions.CookieStore) (botup.Plugin, error) {
	session, _ := store.Get(r, "login")

	form, err := Validate(r, AddPluginRules)
	if err != nil {
		return botup.Plugin{}, ValidationErr
	}

	return botup.Plugin{
		UID:    session.Values["uid"].(string),
		BID:    form["bid"].(int),
		Plugin: form["plugin"].(string),
	}, nil
}

func RemoveBot(r *http.Request, store *sessions.CookieStore) (botup.Bot, error) {
	session, _ := store.Get(r, "login")

	form, err := Validate(r, RemoveBotRules)
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

	form, err := Validate(r, RemovePluginRules)
	if err != nil {
		return botup.Plugin{}, ValidationErr
	}

	return botup.Plugin{
		UID:    session.Values["uid"].(string),
		BID:    form["bid"].(int),
		Plugin: form["plugin"].(string),
	}, nil
}
