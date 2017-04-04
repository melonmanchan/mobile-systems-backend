package handlers

import (
	"encoding/json"
	"net/http"

	"../app"
	"../types"
	"github.com/gorilla/mux"
)

// SubjectHandler ...
func SubjectHandler(app app.App, r *mux.Router) {
	client := app.Client

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		subjects := client.GetSubjects()
		APIResp := types.APIResponse{Result: subjects, Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("GET")
}
