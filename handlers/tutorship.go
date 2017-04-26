package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/melonmanchan/mobile-systems-backend/app"
	"github.com/melonmanchan/mobile-systems-backend/models"
	"github.com/melonmanchan/mobile-systems-backend/types"
	"github.com/melonmanchan/mobile-systems-backend/utils"

	"github.com/gorilla/mux"
	"github.com/maddevsio/fcm"
)

// TutorshipHandler ...
func TutorshipHandler(app app.App, r *mux.Router) {
	client := app.Client
	firebase := app.Firebase

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(types.UserKey).(*models.User)
		decoder := json.NewDecoder(r.Body)

		var req types.CreateTutorShipRequest
		defer r.Body.Close()

		err := decoder.Decode(&req)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGenericRead}, http.StatusBadRequest)
			return
		}

		err = client.CreateTutorship(req.TutorID, user.ID)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorCreatingTutorship}, http.StatusBadRequest)
			return
		}

		APIResp := types.APIResponse{Status: 201}
		encoded, err := json.Marshal(APIResp)
		w.Write(encoded)

		tutor, _ := client.GetUserByID(req.TutorID)

		firebase.SendNotification(tutor.DeviceTokens, fcm.Notification{
			Title: "New tutee!",
			Body:  fmt.Sprintf("You have a new tutee %s %s", user.FirstName, user.LastName),
		})

	}).Methods("POST")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(types.UserKey).(*models.User)

		tutors, err := client.GetUserTutors(user)

		if err != nil {
			log.Println(err)
			utils.FailResponse(w, []types.APIError{types.ErrorGetTutorships}, http.StatusBadRequest)
			return
		}

		tutees, err := client.GetUserTutees(user)

		if err != nil {
			log.Println(err)
			utils.FailResponse(w, []types.APIError{types.ErrorGetTutorships}, http.StatusBadRequest)
			return
		}

		resp := types.TutorshipsResponse{Tutors: tutors, Tutees: tutees}
		APIResp := types.APIResponse{Result: resp, Status: 200}
		encoded, err := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("GET")
}
