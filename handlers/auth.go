package handlers

import (
	"net/http"

	"../models"
	"github.com/gorilla/mux"
)

// AuthHandler ...
func AuthHandler(client *models.Client, r *mux.Router) {

	// Logging in
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		str, _ := client.GetUser()

		w.Write([]byte(*str))
	})

}
