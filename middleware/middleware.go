package middleware

import (
	"net/http"

	"../types"
	"../utils"

	"github.com/urfave/negroni"
)

// SetContentType ...
func SetContentType() negroni.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		next(rw, r)
	}
}

// JSONRecovery ...
func JSONRecovery() negroni.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		defer func() {
			if err := recover(); err != nil {
				utils.FailResponse(rw, []types.APIError{types.ErrorGenericServer}, http.StatusInternalServerError)
			}
		}()
		next(rw, r)
	}
}

// NotFoundHandler ...
func NotFoundHandler(rw http.ResponseWriter, r *http.Request) {
	utils.FailResponse(rw, []types.APIError{types.ErrorGenericNotFound}, http.StatusNotFound)
}
