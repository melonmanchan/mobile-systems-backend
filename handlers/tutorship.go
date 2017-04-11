package handlers

import (
	"net/http"

	"../app"

	"github.com/gorilla/mux"
)

// TutorshipHandler ...
func TutorshipHandler(app app.App, r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("POST")
}
