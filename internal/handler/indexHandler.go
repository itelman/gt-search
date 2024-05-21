package handler

import (
	"functions/internal/models"
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
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

	artists, err := models.GetAllArtists()
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, artists)
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}
}
