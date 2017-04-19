package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

		event, err := client.CreateNewFreeEvent(user, req.StartTime)

		log.Println(err)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorCreateEvent}, http.StatusBadRequest)
			return
		}

		APIResp := types.APIResponse{Result: event, Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("POST")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(types.UserKey).(*models.User)

		decoder := json.NewDecoder(r.Body)

		var req models.Event
		defer r.Body.Close()

		err := decoder.Decode(&req)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGenericRead}, http.StatusBadRequest)
			return
		}

		err = client.RemoveTime(user, &req)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorDeleteEvent}, http.StatusBadRequest)
			return
		}

		APIResp := types.APIResponse{Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("DELETE")

	r.HandleFunc("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tutorID, _ := strconv.ParseInt(vars["id"], 10, 64)

		events, err := client.GetTutorFreeTimes(tutorID)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorFreeTimeGet}, http.StatusBadRequest)
			return
		}

		APIResp := types.APIResponse{Result: events, Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(types.UserKey).(*models.User)

		var resp types.GetEventsResponse

		ownTimes, err := client.GetTutorOwnTimes(user)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorFreeTimeGet}, http.StatusBadRequest)
			return
		}

		log.Println(err)

		tuteeTimes, err := client.GetTuteeTimes(user)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorFreeTimeGet}, http.StatusBadRequest)
			return
		}

		log.Println(err)

		resp.OwnEvents = ownTimes
		resp.ReservedEvents = tuteeTimes

		APIResp := types.APIResponse{Result: resp, Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("GET")
}
