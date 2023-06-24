package routers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func GetBaseRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Two Quotes"))
	})

	r.Mount("/objects", getObjectRouter())
	r.Mount("/auth", getAuthRouter())
	r.Mount("/user", getUserRouter())

	return r
}
