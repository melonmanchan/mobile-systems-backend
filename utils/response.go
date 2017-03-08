package utils

import (
	"encoding/json"
	"net/http"

	"../types"
)

// FailResponse ...
func FailResponse(w http.ResponseWriter, errs []error, status int) {
	var errorMessages []string

	for _, e := range errs {
		errorMessages = append(errorMessages, e.Error())
	}

	resp := types.APIResponse{Errors: errorMessages, Status: status}

	errJSON, _ := json.Marshal(resp)

	w.Write(errJSON)
}
