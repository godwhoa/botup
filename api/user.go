package api

import (
	"errors"
	"github.com/godwhoa/random-shit/botup.me/botup"
	"github.com/godwhoa/random-shit/crypt"
	"github.com/gorilla/sessions"
	"net/http"
)

type User struct {
	service *botup.UserService
	store   *sessions.CookieStore
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
		w.Write([]byte("fields_missing"))
		return
	}
	// Hash password and generate UID
	user.Pass = crypt.Hash(user.Pass)
	user.UID = crypt.UID(user.User)

	// Insert into db
	err = u.service.CreateUser(user)
	// Handle taken user/email
	if err == botup.UserAlreadyExists {
		w.Write([]byte("user_taken"))

	} else if err == nil {
		w.Write([]byte("user_created"))
	} else {
		// TODO: log error
		w.Write([]byte("something_went_wrong"))
	}
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	// Form parse/validation
	email := r.FormValue("email")
	pass := r.FormValue("pass")
	if email == "" || pass == "" {
		w.Write([]byte("fields_missing"))
		return
	}
	// Query db
	user, err := u.service.GetUser(email)
	// Handle wrong email/password
	if err == botup.UserDoesNotExist {
		w.Write([]byte("no_exist"))
		return
	} else if err != nil {
		w.Write([]byte("something_went_wrong"))
		return
	}
	// Set-session
	session, _ := u.store.Get(r, "login")
	session.Values["uid"] = user.UID
	session.Save(r, w)
	w.Write([]byte("logged_in"))
}
