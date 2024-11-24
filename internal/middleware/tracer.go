package middleware

import (
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

// Middleware for generating a request trace ID.
func Tracer(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Create a traceID for this request.
			// Set the traceID in the header to pass to other services.
			traceID := r.Header.Get("X-Trace-Id")
			if traceID == "" {
				r.Header.Set("X-Trace-Id", traceID)
				traceID = uuid.New().String()
			}

			// Create a logger to pass into the request context.
			logger := zerolog.New(os.Stdout).With().
				Timestamp().
				Str("traceId", traceID).
				Str("host", r.Host).
				Str("method", r.Method).
				Str("path", r.URL.Path).
				Logger()

			next.ServeHTTP(w, r.WithContext(logger.WithContext(r.Context())))
		},
	)
}
