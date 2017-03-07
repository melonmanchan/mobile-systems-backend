package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"../app"
	"../models"
	"../utils"
	"github.com/gorilla/mux"
)

// AuthHandler ...
func AuthHandler(app app.App, r *mux.Router) {
	client := app.Client
	config := app.Config

	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var req RegisterRequest
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

		user := req.ToUser()

		err = client.CreateUser(&user)

		if err != nil {
			panic(err)
		}

		token, expiresAt, err := utils.CreateUserToken(user, config)

		if err != nil {
			panic(err)
		}

		resp.ExpiresAt = expiresAt
		resp.CreatedAt = time.Now()
		resp.Token = token
		resp.User = &user

		APIResp := APIResponse{Result: resp}
		encoded, err := json.Marshal(APIResp)
		w.Write(encoded)

	}).Methods("POST")

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

		user, err := client.GetUserByEmail(req.Email, models.NormalAuth)

		if err != nil {
			panic(err)
		}

		err = user.IsPasswordValid(req.Password)

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			panic(err)
		}

		token, expiresAt, err := utils.CreateUserToken(*user, config)

		if err != nil {
			panic(err)
		}

		resp.ExpiresAt = expiresAt
		resp.CreatedAt = time.Now()
		resp.Token = token
		resp.User = user

		APIResp := APIResponse{Result: resp}
		encoded, err := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("POST")
}
