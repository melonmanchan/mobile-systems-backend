package handlers

import (
	"net/http"

	"../app"

	"github.com/gorilla/mux"
)

// EventHandler ...
func EventHandler(app app.App, r *mux.Router) {

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("POST")
}
