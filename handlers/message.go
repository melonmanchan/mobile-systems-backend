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

//MessageHandler ...
func MessageHandler(app app.App, r *mux.Router) {
	client := app.Client

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("POST")

	r.HandleFunc("/latest", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(types.UserKey).(*models.User)

		messages, err := client.GetUserLatestReceivedMessages(user)

		log.Println(err)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGetLatest}, http.StatusBadRequest)
			return
		}

		APIResp := types.APIResponse{Result: messages, Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("GET")
}
