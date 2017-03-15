package handlers

import (
	"log"
	"net/http"

	"../app"
	"../models"
	"../types"
	"github.com/gorilla/mux"
)

// UserHandler ...
func UserHandler(app app.App, r *mux.Router) {
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(types.UserKey).(*models.User)
		log.Print(user)
	}).Methods("POST")
}
