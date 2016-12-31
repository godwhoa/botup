package api

import (
	"errors"
	"github.com/godwhoa/random-shit/botup.me/botup"
	"github.com/godwhoa/random-shit/botup.me/crypt"
	"github.com/gorilla/sessions"
	"net/http"
)

type User struct {
	Service botup.UserService
	Store   *sessions.CookieStore
}

var invalidForm = errors.New("Some form fields are missing")

// Validates registration form
// Puts data inside User struct
func validate_reg_form(r *http.Request) (botup.User, error) {
	email := r.FormValue("email")
	user := r.FormValue("user")
	pass := r.FormValue("pass")
	if email == "" || user == "" || pass == "" {
		return botup.User{}, invalidForm
	}
	return botup.User{Email: email, User: user, Pass: pass}, nil
}

func (u *User) Register(w http.ResponseWriter, r *http.Request) {
	// Parse/validate form
	user, err := validate_reg_form(r)
	if err != nil {
		w.Write(botup.ERR_FIELDS_MISSING)
		return
	}
	// Hash password and generate UID
	user.Pass, _ = crypt.Hash(user.Pass)
	user.UID, _ = crypt.UID(user.User)

	// Insert into db
	err = u.Service.CreateUser(user)
	// Handle taken user/email
	if err == botup.UserAlreadyExists {
		w.Write(botup.ERR_USER_TAKEN)

	} else if err == nil {
		w.Write(botup.OK_USER_CREATED)
	} else {
		// TODO: log error
		w.Write(botup.ERR_INTERNAL)
	}
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	session, _ := u.Store.Get(r, "login")
	// Form parse/validation
	email := r.FormValue("email")
	pass := r.FormValue("pass")
	if email == "" || pass == "" {
		w.Write(botup.ERR_FIELDS_MISSING)
		return
	}
	// Query db
	user, err := u.Service.GetUser(email)
	// Handle wrong email/password
	if err == botup.UserDoesNotExist {
		w.Write(botup.ERR_WRONG_CREDENTIALS)
		return
	} else if err != nil {
		w.Write(botup.ERR_INTERNAL)
		return
	}
	// Set-session
	session.Values["uid"] = user.UID
	session.Save(r, w)
	w.Write(botup.OK_LOGGED_IN)
}

func (u *User) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := u.Store.Get(r, "login")
	if err != nil {
		w.Write(botup.ERR_NOT_LOGGED_IN)
		return
	}
	session.Values["uid"] = nil
	session.Save(r, w)
	w.Write(botup.OK_LOGGED_OUT)
}
