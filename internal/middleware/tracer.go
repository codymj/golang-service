package middleware

import (
	"net/http"

	"github.com/google/uuid"
)

// Middleware for generating a request trace ID.
func Tracer(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("X-Trace-ID", uuid.New().String())

		next.ServeHTTP(w, r)
	})
}
