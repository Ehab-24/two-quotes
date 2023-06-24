package routers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"suraj.com/refine/data"
)

func getUserRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", getCurrentUser)
	r.Delete("/", deleteAllUsers)

	return r
}

func deleteAllUsers(w http.ResponseWriter, r *http.Request) {
	res, err := data.UserDeleteAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
