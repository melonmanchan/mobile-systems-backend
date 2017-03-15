package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"../app"
	"../models"
	"../types"
	"../utils"
	"github.com/gorilla/mux"
)

// UserHandler ...
func UserHandler(app app.App, r *mux.Router) {
	client := app.Client

	r.HandleFunc("/register_device", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(types.UserKey).(*models.User)
		decoder := json.NewDecoder(r.Body)

		var req types.DeviceRegisterRequest
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

		err = client.AddTokenToUser(user, req.Token)

		log.Print(err)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGenericRead}, http.StatusBadRequest)
			return
		}

		APIResp := types.APIResponse{Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("POST")
}
