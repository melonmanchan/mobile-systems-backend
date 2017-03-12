package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"../app"
	"../models"
	"../types"
	"../utils"
	"github.com/gorilla/mux"
)

// AuthHandler ...
func AuthHandler(app app.App, r *mux.Router) {
	client := app.Client
	config := app.Config

	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var req types.RegisterRequest
		var resp types.LoginResponse

		defer r.Body.Close()

		err := decoder.Decode(&req)

		if err != nil {
			panic(err)
		}

		valid, errs := req.IsValid()

		if !valid {
			utils.FailResponse(w, errs, http.StatusBadRequest)
			return
		}

		user := req.ToUser()

		err = client.CreateUser(&user)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorRegisterAlreadyExists}, http.StatusInternalServerError)
			return
		}

		token, expiresAt, err := utils.CreateUserToken(user, config)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGenericTokenCreate}, http.StatusInternalServerError)
			return
		}

		resp.ExpiresAt = expiresAt
		resp.CreatedAt = time.Now()
		resp.Token = token
		resp.User = &user

		APIResp := types.APIResponse{Result: resp, Status: 201}
		encoded, err := json.Marshal(APIResp)
		w.Write(encoded)

	}).Methods("POST")

	// Logging in
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var req types.LoginRequest
		var resp types.LoginResponse

		defer r.Body.Close()

		err := decoder.Decode(&req)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGenericRead}, http.StatusBadRequest)
			return
		}

		valid, errs := req.IsValid()

		if !valid {
			utils.FailResponse(w, errs, http.StatusBadRequest)
			return
		}

		user, err := client.GetUserByEmail(req.Email, models.NormalAuth)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorLoginUserNotFound}, http.StatusBadRequest)
			return
		}

		err = user.IsPasswordValid(req.Password)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorLoginWrongPassword}, http.StatusBadRequest)
			return
		}

		token, expiresAt, err := utils.CreateUserToken(*user, config)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGenericTokenCreate}, http.StatusBadRequest)
			return
		}

		resp.ExpiresAt = expiresAt
		resp.CreatedAt = time.Now()
		resp.Token = token
		resp.User = user

		APIResp := types.APIResponse{Result: resp, Status: 200}
		encoded, err := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("POST")
}
