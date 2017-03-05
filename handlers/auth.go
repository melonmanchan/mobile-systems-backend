package handlers

import (
	"encoding/json"
	"net/http"

	"../app"
	"github.com/gorilla/mux"
)

// AuthHandler ...
func AuthHandler(app app.App, r *mux.Router) {
	client := app.Client

	// Logging in
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var req LoginRequest
		var resp LoginResponse

		defer r.Body.Close()

		err := decoder.Decode(&req)

		if err != nil {
			panic(err)
		}

		valid, errs := req.IsValid()

		if !valid {
			w.WriteHeader(http.StatusBadRequest)
			panic(errs)
		}

		user, err := client.GetUserByEmail(req.Email)

		if err != nil {
			panic(err)
		}

		encoded, err := json.Marshal(user)
		w.Write(encoded)
	}).Methods("POST")
}
