package routers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"suraj.com/refine/data"
)

func getCommentsRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/{uid}", getCommentsByUser)

	return r
}

func getCommentsByPostRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", getCommentsByPost)

	return r
}

func getCommentsByPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	commentDoc, err := data.CommentsFindByPostId(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(commentDoc.Comments)
}

func getCommentsByUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	comments, err := data.CommentsFindByUserId(chi.URLParam(r, "uid"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comments)
}
