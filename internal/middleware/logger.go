package middleware

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

// Middleware for logging request information.
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request info.
		start := time.Now()
		log.Info().
			Str("traceID", r.Header.Get("X-Trace-ID")).
			Str("host", r.Host).
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Msg("init")

		// Call next handler in the chain.
		next.ServeHTTP(w, r)

		// Log request info.
		log.Info().
			Str("traceID", r.Header.Get("X-Trace-ID")).
			Str("host", r.Host).
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Dur("duration", time.Duration(time.Since(start).Nanoseconds())).
			Msg("done")
	})
}
