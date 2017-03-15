package handlers

import (
	"encoding/json"
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
	config := app.Config

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

		deviceAlreadyRegistered := false

		for _, token := range user.DeviceTokens {
			if token == req.Token {
				deviceAlreadyRegistered = true
				break
			}
		}

		if deviceAlreadyRegistered == false {
			err = client.AddTokenToUser(user, req.Token)
		}

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGenericRead}, http.StatusBadRequest)
			return
		}

		//utils.SendNotification(config, user.DeviceTokens, fcm.Notification{
		//Title: "Hello",
		//Body:  "World",
		//})

		APIResp := types.APIResponse{Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("POST")
}
