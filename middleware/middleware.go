package middleware

import (
	"net/http"

	"github.com/urfave/negroni"
)

// SetContentType ...
func SetContentType() negroni.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		next(rw, r)
	}
}
