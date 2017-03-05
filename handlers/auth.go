package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"../app"
	"github.com/gorilla/mux"
)

// AuthHandler ...
func AuthHandler(app app.App, r *mux.Router) {
	client := app.Client

	// Logging in
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		user, err := client.GetUserByEmail("a@a.com")

		if err != nil {
			log.Fatal(err)
		}

		log.Println(user.ID)
		encoded, err := json.Marshal(user)
		w.Write(encoded)
	}).Methods("POST")
}
