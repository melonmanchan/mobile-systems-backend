package handlers

import (
	"log"
	"net/http"

	"../models"
	"github.com/gorilla/mux"
)

// AuthHandler ...
func AuthHandler(client *models.Client, r *mux.Router) {

	// Logging in
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		user, err := client.GetUserByEmail("a@a.com")

		if err != nil {
			log.Fatal(err)
		}

		log.Print(user)

		w.Write([]byte(user.FirstName))
	})

}
