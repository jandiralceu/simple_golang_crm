package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	println("Hello, World!")

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

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
