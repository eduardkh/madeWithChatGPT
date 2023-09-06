package server

import (
	"fmt"
	"net/http"

	"myapp/internal/vault"

	"github.com/gorilla/sessions"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "Welcome to the main page. Please <a href='/login'>log in</a> to access the secret page.")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<form action='/login' method='POST'>Username: <input type='text' name='username'><br>Password: <input type='password' name='password'><br><input type='submit' value='Log in'></form>")
}

func loginPostHandler(vaultClient *vault.Client, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		token, err := vaultClient.AuthenticateWithUserPass(username, password)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		session, _ := r.Context().Value("session").(*sessions.Session)
		session.Values["token"] = token
		session.Save(r, w)

		http.Redirect(w, r, "/secret", http.StatusFound)
	}
}

func secretHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "Welcome to the secret page!")
}
