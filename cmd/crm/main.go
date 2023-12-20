package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Route("/api/v1/customers", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			// Code goes here
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			// Code goes here
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			// Code goes here
		})
		r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
			// Code goes here
		})
		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			// Code goes here
		})
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
