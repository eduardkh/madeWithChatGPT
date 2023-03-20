package server

import (
	"context"
	"net/http"

	"myapp/internal/vault"

	"github.com/gorilla/sessions"
)

const SessionName = "myapp_session"

func SessionMiddleware(store *sessions.CookieStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, _ := store.Get(r, SessionName)
			ctx := r.Context()
			ctx = context.WithValue(ctx, "session", session)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AuthMiddleware(vaultClient *vault.Client) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, _ := r.Context().Value("session").(*sessions.Session)
			token, ok := session.Values["token"].(string)
			if !ok {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}

			err := vaultClient.Authorize(token, "your-secret-path") // Replace "your-secret-path" with the actual secret path in Vault
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
