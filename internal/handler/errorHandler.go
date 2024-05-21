package handler

import (
	"functions/internal/models"
	"net/http"
	"text/template"
)

func ErrorPageHandler(w http.ResponseWriter, statusCode int) {
	errTmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	err = errTmpl.Execute(w, models.ErrorConstructor(statusCode))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
