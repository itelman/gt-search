package server

import (
	"fmt"
	"functions/internal/handler"
	"log"
	"net/http"
)

func RunServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.IndexHandler)
	mux.HandleFunc("/results", handler.ResultsHandler)
	mux.HandleFunc("/artists", handler.ArtistHandler)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Server is running on http://localhost:8080/...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
