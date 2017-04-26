package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/melonmanchan/mobile-systems-backend/app"
	"github.com/melonmanchan/mobile-systems-backend/models"
	"github.com/melonmanchan/mobile-systems-backend/types"
	"github.com/melonmanchan/mobile-systems-backend/utils"

	"github.com/gorilla/mux"
)

//MessageHandler ...
func MessageHandler(app app.App, r *mux.Router) {
	client := app.Client
	firebase := app.Firebase

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(types.UserKey).(*models.User)

		decoder := json.NewDecoder(r.Body)

		var req types.CreateMessageRequest
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

		msg, err := client.CreateMessage(user.ID, req.Receiver, req.Content)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorCreateMessage}, http.StatusBadRequest)
			return
		}

		APIResp := types.APIResponse{Result: msg, Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)

		receiver, err := client.GetUserByID(req.Receiver)
		err = firebase.SendMessage(receiver.DeviceTokens, msg)
	}).Methods("POST")

	r.HandleFunc("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		user := r.Context().Value(types.UserKey).(*models.User)
		recipientID, _ := strconv.ParseInt(vars["id"], 10, 64)

		toStr := r.URL.Query().Get("to")
		fromStr := r.URL.Query().Get("from")

		toInt, _ := strconv.ParseInt(toStr, 10, 64)
		fromInt, _ := strconv.ParseInt(fromStr, 10, 64)

		messages, err := client.GetConversation(user.ID, recipientID, fromInt, toInt)

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
