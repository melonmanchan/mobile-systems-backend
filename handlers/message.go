package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../app"
	"../models"
	"../types"
	"../utils"

	"github.com/gorilla/mux"
)

//MessageHandler ...
func MessageHandler(app app.App, r *mux.Router) {
	client := app.Client

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(types.UserKey).(*models.User)

		decoder := json.NewDecoder(r.Body)

		var req types.Create
		defer r.Body.Close()

		err := decoder.Decode(&req)

		APIResp := types.APIResponse{Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("POST")

	r.HandleFunc("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		user := r.Context().Value(types.UserKey).(*models.User)
		recipientID, _ := strconv.ParseInt(vars["id"], 10, 64)

		messages, err := client.GetConversation(user.ID, recipientID)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGetLatest}, http.StatusBadRequest)
			return
		}

		APIResp := types.APIResponse{Result: messages, Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("GET")

	r.HandleFunc("/latest", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(types.UserKey).(*models.User)

		messages, err := client.GetUserLatestReceivedMessages(user)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGetLatest}, http.StatusBadRequest)
			return
		}

		APIResp := types.APIResponse{Result: messages, Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("GET")
}
