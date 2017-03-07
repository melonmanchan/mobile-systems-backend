package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../handlers"

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
				var errorMessages []string

				switch e := err.(type) {
				case error:
					errorMessages = append(errorMessages, e.Error())
				case []error:
					for _, e := range e {
						errorMessages = append(errorMessages, e.Error())
					}
				}

				resp := handlers.APIResponse{Errors: errorMessages}
				errJSON, _ := json.Marshal(resp)

				rw.Write(errJSON)
			}
		}()
		next(rw, r)
	}
}

// NotFoundHandler ...
func NotFoundHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusNotFound)
	panic(fmt.Errorf("Matching route not found for %s with method %s", r.RequestURI, r.Method))
}
