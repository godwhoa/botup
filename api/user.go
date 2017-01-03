package api

import (
	"net/http"

	"github.com/godwhoa/random-shit/botup.me/botup"
	"github.com/godwhoa/random-shit/botup.me/crypt"
	"github.com/gorilla/sessions"
)

type User struct {
	Service    botup.UserService
	Store      *sessions.CookieStore
	LoginCache map[string]string
}

func (u *User) Register(w http.ResponseWriter, r *http.Request) {
	user, err := validate_reg_form(r)
	if err != nil {
		w.Write(botup.ERR_FIELDS_MISSING)
		return
	}

	// Hash password and generate UID
	user.Pass, _ = crypt.Hash(user.Pass)
	user.UID, _ = crypt.UID(user.User)

	err = u.Service.CreateUser(user)
	switch err {
	case botup.UserAlreadyExists:
		w.Write(botup.ERR_USER_TAKEN)
	case nil:
		w.Write(botup.OK_USER_CREATED)
	default:
		// TODO: log error
		w.Write(botup.ERR_INTERNAL)
	}
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	session, _ := u.Store.Get(r, "login")

	email, pass := r.FormValue("email"), r.FormValue("pass")
	if email == "" || pass == "" {
		w.Write(botup.ERR_FIELDS_MISSING)
		return
	}

	user, err := u.Service.GetUser(email)
	switch err {
	case botup.UserDoesNotExist:
		w.Write(botup.ERR_WRONG_CREDENTIALS)
	case nil:
		if crypt.Verify(user.Pass, pass) {
			session.Values["uid"] = user.UID
			u.LoginCache[user.UID] = "loggedin"
			w.Write(botup.OK_LOGGED_IN)
		} else {
			w.Write(botup.ERR_WRONG_CREDENTIALS)
		}
	default:
		// TODO: log error
		w.Write(botup.ERR_INTERNAL)
	}
	session.Save(r, w)
}

func (u *User) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := u.Store.Get(r, "login")
	if err != nil {
		w.Write(botup.ERR_NOT_LOGGED_IN)
		return
	}

	u.LoginCache[session.Values["uid"].(string)] = "loggedout"
	delete(session.Values, "uid")
	session.Save(r, w)

	w.Write(botup.OK_LOGGED_OUT)
}
