package middleware

import (
	"context"
	"net/http"
	"strings"

	"../app"
	"../types"
	"../utils"

	"github.com/urfave/negroni"
)

// CreateResolveUser ...
func CreateResolveUser(app app.App) func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	config := app.Config

	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		token := r.Header.Get("Authorization")

		if token == "" {
			utils.FailResponse(rw, []types.APIError{types.ErrorGenericTokenMissing}, http.StatusForbidden)
			return
		}

		tokenSansBearer := strings.TrimPrefix(token, "Bearer ")

		user, err := utils.DecodeUserFromJWT(tokenSansBearer, config)

		if err != nil {
			utils.FailResponse(rw, []types.APIError{types.ErrorGenericTokenInvalid}, http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), types.UserKey, user)
		r = r.WithContext(ctx)

		next(rw, r)
	}
}

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
