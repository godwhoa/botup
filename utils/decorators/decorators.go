package decorators

import (
	"github.com/godwhoa/random-shit/botup.me/botup"
	"github.com/gorilla/sessions"
	"net/http"
)

func Auth(fn func(w http.ResponseWriter, r *http.Request), store *sessions.CookieStore, cache *botup.CacheService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := u.Store.Get(r, "login")
		if err != nil || session.Values["uid"] == nil {
			w.Write(botup.ERR_NOT_LOGGED_IN)
			return
		}

		status, _ := cache.Get(session.Values["uid"])
		if status == "loggedin" {
			fn(w, r)
		} else {
			w.Write(botup.ERR_NOT_LOGGED_IN)
		}
	}
}
