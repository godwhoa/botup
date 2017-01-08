package api

import (
	"github.com/godwhoa/random-shit/botup.me/botup"
	"github.com/godwhoa/random-shit/botup.me/crypt"
	"github.com/gorilla/sessions"
	"github.com/uber-go/zap"
	"net/http"
)

type User struct {
	Service    botup.UserService
	Store      *sessions.CookieStore
	LoginCache map[string]string
	Log        zap.Logger
}

func (u *User) Register(w http.ResponseWriter, r *http.Request) {
	user, err := validate_reg_form(r, w)
	if err != nil {
		w.Write(botup.ERR_FIELDS_MISSING)
		u.Log.Error("api.user.register",
			zap.Error(err),
			zap.String("info", "register form fields missing"),
			zap.String("ip", r.RemoteAddr),
		)
		return
	}

	// Hash password and generate UID
	user.Pass, _ = crypt.Hash(user.Pass)
	user.UID, _ = crypt.UID(user.User)

	err = u.Service.CreateUser(user)
	switch err {
	case botup.UserAlreadyExists:
		w.Write(botup.ERR_USER_TAKEN)
		u.Log.Error("api.user.register",
			zap.Error(err),
			zap.String("info", "user taken"),
			zap.String("user", user.User),
			zap.String("email", user.Email),
			zap.String("ip", r.RemoteAddr),
		)
	case nil:
		w.Write(botup.OK_USER_CREATED)
		u.Log.Info("api.user.register",
			zap.String("info", "user created"),
			zap.String("user", user.User),
			zap.String("email", user.Email),
			zap.String("ip", r.RemoteAddr),
		)
	default:
		w.Write(botup.ERR_INTERNAL)
		u.Log.Error("api.user.register",
			zap.Error(err),
			zap.String("info", "error creating user"),
			zap.String("user", user.User),
			zap.String("email", user.Email),
			zap.String("ip", r.RemoteAddr),
		)
	}
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	session, _ := u.Store.Get(r, "login")

	email, pass := r.FormValue("email"), r.FormValue("pass")
	if email == "" || pass == "" {
		w.Write(botup.ERR_FIELDS_MISSING)
		u.Log.Error("api.user.login",
			zap.Error(invalidForm),
			zap.String("info", "login form fields missing"),
			zap.String("ip", r.RemoteAddr),
		)
		return
	}

	user, err := u.Service.GetUser(email)
	switch err {
	case botup.UserDoesNotExist:
		w.Write(botup.ERR_WRONG_CREDENTIALS)
		u.Log.Info("api.user.login",
			zap.Error(err),
			zap.String("info", "wrong credentials"),
			zap.String("email", email),
			zap.String("pass", pass),
			zap.String("ip", r.RemoteAddr),
		)
	case nil:
		if crypt.Verify(user.Pass, pass) {
			session.Values["uid"] = user.UID
			session.Save(r, w)

			u.LoginCache[user.UID] = "loggedin"
			w.Write(botup.OK_LOGGED_IN)

			u.Log.Info("api.user.login",
				zap.String("info", "user logged in"),
				zap.String("email", user.Email),
				zap.String("user", user.User),
				zap.String("uid", user.UID),
				zap.String("ip", r.RemoteAddr),
			)
		} else {
			w.Write(botup.ERR_WRONG_CREDENTIALS)
			u.Log.Error("api.user.login",
				zap.String("info", "wrong credentials at nil"),
				zap.String("email", email),
				zap.String("pass", pass),
				zap.String("ip", r.RemoteAddr),
			)
		}
	default:
		w.Write(botup.ERR_INTERNAL)
		u.Log.Error("api.user.login",
			zap.Error(err),
			zap.String("info", "error logging in"),
			zap.String("user", user.User),
			zap.String("email", user.Email),
			zap.String("ip", r.RemoteAddr),
		)
	}
	session.Save(r, w)
}

func (u *User) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := u.Store.Get(r, "login")
	uid, ok := session.Values["uid"]
	if err != nil || !ok {
		w.Write(botup.ERR_NOT_LOGGED_IN)
		u.Log.Error("api.user.logout",
			zap.Error(err),
			zap.String("info", "error logging out"),
			zap.String("ip", r.RemoteAddr),
		)
		return
	}
	u.LoginCache[uid.(string)] = "loggedout"
	delete(session.Values, "uid")
	session.Save(r, w)

	w.Write(botup.OK_LOGGED_OUT)
	u.Log.Info("api.user.logout",
		zap.String("info", "user logged out"),
		zap.String("uid", uid.(string)),
		zap.String("ip", r.RemoteAddr),
	)
}
