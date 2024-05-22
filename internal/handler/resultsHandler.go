package handler

import (
	"encoding/json"
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

	tmpl, err := template.ParseFiles("templates/results.html")
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

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/results/json" {
		ErrorPageHandler(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorPageHandler(w, http.StatusMethodNotAllowed)
		return
	}

	results, err := models.SetResults(r.FormValue("q"))
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(results)
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
