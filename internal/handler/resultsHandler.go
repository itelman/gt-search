package handler

import (
	"functions/internal/models"
	"net/http"
	"text/template"
)

func ResultsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/results" {
		ErrorPageHandler(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorPageHandler(w, http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}

	results, err := models.SetResults(r.FormValue("q"))
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, results)
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}
}
