package utils

import (
	"encoding/json"
	"net/http"

	"github.com/melonmanchan/mobile-systems-backend/types"
)

// FailResponse ...
func FailResponse(w http.ResponseWriter, errs []types.APIError, status int) {
	resp := types.APIResponse{Errors: errs, Status: status}
	errJSON, _ := json.Marshal(resp)
	w.Write(errJSON)
}
