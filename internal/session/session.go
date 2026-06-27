package session

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

const (
	sessionName = "fc_session"
	userIDKey   = "user_id"
)

var store *sessions.CookieStore

func Init() {
	secret := os.Getenv("SESSION_SECRET")
	if secret == "" {
		secret = "dev-secret-change-in-production!!"
	}
	store = sessions.NewCookieStore([]byte(secret))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
}

func SetUserID(w http.ResponseWriter, r *http.Request, id string) error {
	s, _ := store.Get(r, sessionName)
	s.Values[userIDKey] = id
	return s.Save(r, w)
}

func GetUserID(r *http.Request) string {
	s, err := store.Get(r, sessionName)
	if err != nil {
		return ""
	}
	id, _ := s.Values[userIDKey].(string)
	return id
}

func Clear(w http.ResponseWriter, r *http.Request) {
	s, _ := store.Get(r, sessionName)
	s.Options.MaxAge = -1
	s.Save(r, w) //nolint:errcheck
}

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if GetUserID(r) == "" {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
