package handlers

import (
	"net/http"
	"text/template"
)

// MainPageHandler handles the main project page. It contains information about the project.
func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("internal/templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := page.Execute(w, nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
