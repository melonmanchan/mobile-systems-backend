package utils

import (
	"encoding/json"
	"net/http"

	"../types"
)

// FailResponse ...
func FailResponse(w http.ResponseWriter, errs []types.APIError, status int) {
	resp := types.APIResponse{Errors: errs, Status: status}
	errJSON, _ := json.Marshal(resp)
	w.Write(errJSON)
}
