package handlers

import (
	"net/http"

	"../app"

	"github.com/gorilla/mux"
)

//MessageHandler ...
func MessageHandler(app app.App, r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("POST")
}
