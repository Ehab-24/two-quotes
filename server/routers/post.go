package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"suraj.com/refine/data"
	"suraj.com/refine/models"
)

func getPostRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/", getPosts)
		r.Post("/", createPost)
		r.Delete("/", deletePosts)
	})

	r.Route("/{pid}", func(r chi.Router) {
		r.Get("/", getPost)
		r.Delete("/", deletePost)
		// TODO: r.Put("/", updatePost)
		r.Mount("/comments", getCommentsByPostRouter())
	})

	return r
}

func getPost(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "pid")

	res, err := data.PostFindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(res)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "pid")

	res, err := data.PostDeleteById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	res, err := data.PostFindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(res)
}

func deletePosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := data.PostDeleteAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var obj models.Post
	obj.FromJSON(&r.Body)

	// TODO: validate `obj`

	res, err := data.PostCreateOne(&obj)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&res)
}
