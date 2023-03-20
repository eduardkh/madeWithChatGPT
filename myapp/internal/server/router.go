package server

import (
	"myapp/internal/vault"

	"github.com/go-chi/chi"
	"github.com/gorilla/sessions"
)

func NewRouter(store *sessions.CookieStore, vaultClient *vault.Client) *chi.Mux {
	router := chi.NewRouter()

	router.Use(SessionMiddleware(store))

	router.Get("/", indexHandler)
	router.Get("/login", loginHandler)
	router.Post("/login", loginPostHandler(vaultClient, store))

	router.Group(func(r chi.Router) {
		r.Use(AuthMiddleware(vaultClient))
		r.Get("/secret", secretHandler)
	})

	return router
}
