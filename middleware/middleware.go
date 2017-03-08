package middleware

import (
	"fmt"
	"net/http"

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
				var errs []error

				switch e := err.(type) {
				case error:
					errs = append(errs, e)
				case []error:
					errs = e
				}

				utils.FailResponse(rw, errs, http.StatusInternalServerError)
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
