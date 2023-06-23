package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"suraj.com/refine/data"
	"suraj.com/refine/models"
)

func errorResponse(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func GetObjectRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/", getObjects)
		r.Post("/", createObject)
		r.Delete("/", deleteObjects)
	})

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", getObject)
		r.Delete("/", deleteObject)
		// r.Put("/", updateObject)
	})

	return r
}

func getObject(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	res, err := data.ObjectGetById(id)
	if err != nil {
		errorResponse(w, err)
		return
	}

	log.Println(res)
}

func deleteObject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")

	res, err := data.ObjectDeleteById(id)
	if err != nil {
		errorResponse(w, err)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func getObjects(w http.ResponseWriter, r *http.Request) {
	res, err := data.ObjectGetAll()
	if err != nil {
		errorResponse(w, err)
		return
	}

	log.Println(res)
}

func deleteObjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := data.ObjectDeleteAll()
	if err != nil {
		errorResponse(w, err)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func createObject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var obj models.Object
	obj.FromJSON(&r.Body)

	// TODO: validate `obj`

	res, err := data.ObjectCreate(&obj)

	if err != nil {
		errorResponse(w, err)
		return
	}
	json.NewEncoder(w).Encode(&res)
}
