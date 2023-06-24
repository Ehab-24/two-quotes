package routers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"suraj.com/refine/data"
	"suraj.com/refine/models"
)

func getCommentsRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", getCommentsByUser)
	r.Get("/{cid}", getCommentById)

	return r
}

func getCommentsByPostRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", getCommentsByPost)
	r.Post("/", createComment)
	r.Delete("/{cid}", deleteCommentById)

	return r
}

func getCommentsByPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	commentDoc, err := data.CommentFindByPostId(chi.URLParam(r, "pid"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(commentDoc.Comments)
}

func getCommentsByUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	comments, err := data.CommentFindByUserId(r.URL.Query().Get("userId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comments)
}

// ! Whats the point of this?
func getCommentById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// comment, err := daat
}

func deleteCommentById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	res, err := data.CommentDeleteById(chi.URLParam(r, "pid"), chi.URLParam(r, "cid"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var comment models.Comment
	if err := comment.FromJSON(&r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := data.CommentCreateOne(chi.URLParam(r, "pid"), &comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(&res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
