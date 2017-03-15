package handlers

import (
	"log"
	"net/http"

	"../app"
	"github.com/gorilla/mux"
)

// UserHandler ...
func UserHandler(app app.App, r *mux.Router) {
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("test route")
	}).Methods("POST")
}
