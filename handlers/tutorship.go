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

// TutorshipHandler ...
func TutorshipHandler(app app.App, r *mux.Router) {
	client := app.Client

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(types.UserKey).(*models.User)
		decoder := json.NewDecoder(r.Body)

		var req types.CreateTutorShipRequest
		defer r.Body.Close()

		err := decoder.Decode(req)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGenericRead}, http.StatusBadRequest)
			return
		}

		isTutor, err := client.IsUserIDTutor(req.TutorID)

		if !isTutor || err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorNotTutor}, http.StatusBadRequest)
			return
		}

		err = client.CreateTutorship(user.ID, req.TutorID)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorCreatingTutorship}, http.StatusBadRequest)
			return
		}

		APIResp := types.APIResponse{Status: 201}
		encoded, err := json.Marshal(APIResp)
		w.Write(encoded)

	}).Methods("POST")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(types.UserKey).(*models.User)

		tutors, err := client.GetUserTutors(user)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGetTutorships}, http.StatusBadRequest)
			return
		}

		tutees, err := client.GetUserTutees(user)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGetTutorships}, http.StatusBadRequest)
			return
		}

		resp := types.TutorshipsResponse{Tutors: tutors, Tutees: tutees}
		APIResp := types.APIResponse{Result: resp, Status: 200}
		encoded, err := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("GET")
}
