package handler

import (
	"functions/internal/models"
	"net/http"
	"strconv"
	"text/template"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artists" {
		ErrorPageHandler(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorPageHandler(w, http.StatusMethodNotAllowed)
		return
	}

	id_str := r.FormValue("id")

	id, err := strconv.Atoi(id_str)
	if err != nil {
		ErrorPageHandler(w, http.StatusBadRequest)
		return
	}

	artistsList, err := models.GetAllArtists()
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}

	defer func() {
		if recover() != nil {
			ErrorPageHandler(w, http.StatusNotFound)
			return
		}
	}()

	err = artistsList[id-1].GetFullInfo()
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}

	mainTmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}

	err = mainTmpl.Execute(w, artistsList[id-1])
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}
}
