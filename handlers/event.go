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

// EventHandler ...
func EventHandler(app app.App, r *mux.Router) {
	client := app.Client

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(types.UserKey).(*models.User)

		decoder := json.NewDecoder(r.Body)

		var req types.CreateFreeEventRequest
		defer r.Body.Close()

		err := decoder.Decode(&req)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGenericRead}, http.StatusBadRequest)
			return
		}

		event, err := client.CreateNewFreeEvent(user, req.StartTime, req.EndTime)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorCreateEvent}, http.StatusBadRequest)
			return
		}

		APIResp := types.APIResponse{Result: event, Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("POST")
}
