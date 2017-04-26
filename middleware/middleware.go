package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/melonmanchan/mobile-systems-backend/app"
	"github.com/melonmanchan/mobile-systems-backend/types"
	"github.com/melonmanchan/mobile-systems-backend/utils"

	"github.com/urfave/negroni"
)

// CreateResolveUser ...
func CreateResolveUser(app app.App) func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	client := app.Client
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

		fullUser, err := client.GetUserByEmail(user.Email, user.AuthenticationMethod)

		if err != nil {
			log.Print(err)
			utils.FailResponse(rw, []types.APIError{types.ErrorLoginUserNotFound}, http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), types.UserKey, fullUser)
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
				log.Println(err)
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
