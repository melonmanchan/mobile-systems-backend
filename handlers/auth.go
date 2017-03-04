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

		log.Print(user.AuthenticationMethod.ID)
		log.Print(user.AuthenticationMethod.Type)

		log.Print(user.UserType.ID)
		log.Print(user.UserType.Type)

		w.Write([]byte(user.AuthenticationMethod.Type + user.UserType.Type))
	})

}
