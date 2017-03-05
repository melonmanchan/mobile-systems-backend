package middleware

import (
	"encoding/json"
	"fmt"
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

// NotFoundHandler ...
func NotFoundHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusNotFound)

	err := map[string]string{
		"error": fmt.Sprintf("Matching not found for %s with method %s", r.RequestURI, r.Method),
	}

	errJSON, _ := json.Marshal(err)

	rw.Write(errJSON)
}
