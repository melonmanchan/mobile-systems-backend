package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/melonmanchan/mobile-systems-backend/app"
	"github.com/melonmanchan/mobile-systems-backend/models"
	"github.com/melonmanchan/mobile-systems-backend/types"
	"github.com/melonmanchan/mobile-systems-backend/utils"

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
			log.Println(err)
			utils.FailResponse(w, []types.APIError{types.ErrorGenericRead}, http.StatusBadRequest)
			return
		}

		event, err := client.CreateNewFreeEvent(user, req.StartTime)

		log.Println(event)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorCreateEvent}, http.StatusBadRequest)
			return
		}

		APIResp := types.APIResponse{Result: event, Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("POST")

	r.HandleFunc("/remove", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(types.UserKey).(*models.User)

		decoder := json.NewDecoder(r.Body)

		var req models.Event
		defer r.Body.Close()

		err := decoder.Decode(&req)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGenericRead}, http.StatusBadRequest)
			return
		}

		log.Println(req)

		err = client.RemoveTime(user, &req)

		if err != nil {
			log.Println(err)
			utils.FailResponse(w, []types.APIError{types.ErrorDeleteEvent}, http.StatusBadRequest)
			return
		}

		APIResp := types.APIResponse{Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("PUT")

	r.HandleFunc("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tutorID, _ := strconv.ParseInt(vars["id"], 10, 64)

		events, err := client.GetTutorFreeTimes(tutorID)

		if err != nil {
			log.Println(err)
			utils.FailResponse(w, []types.APIError{types.ErrorFreeTimeGet}, http.StatusBadRequest)
			return
		}

		APIResp := types.APIResponse{Result: events, Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("GET")

	r.HandleFunc("/reserve", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(types.UserKey).(*models.User)

		decoder := json.NewDecoder(r.Body)

		var req models.Event
		defer r.Body.Close()

		err := decoder.Decode(&req)
		log.Println(err)
		log.Println(req)
		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGenericRead}, http.StatusBadRequest)
			return
		}

		err = client.ReserveTimeForUser(user, &req)

		log.Println(err)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGenericRead}, http.StatusBadRequest)
			return
		}

		APIResp := types.APIResponse{Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("PUT")

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

		log.Println(ownTimes)
		log.Println(tuteeTimes)

		APIResp := types.APIResponse{Result: resp, Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("GET")
}
