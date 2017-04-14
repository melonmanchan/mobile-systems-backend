package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../app"
	"../types"
	"../utils"
	"github.com/gorilla/mux"
)

// SubjectHandler ...
func SubjectHandler(app app.App, r *mux.Router) {
	client := app.Client

	r.HandleFunc("/{id:[0-9]+}/tutors", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		subjectID, _ := strconv.ParseInt(vars["id"], 10, 64)

		tutors, err := client.GetTutorsBySubjectID(subjectID)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGenericRead}, http.StatusBadRequest)
			return
		}

		APIResp := types.APIResponse{Result: tutors, Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		subjects := client.GetSubjects()
		APIResp := types.APIResponse{Result: subjects, Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("GET")
}
