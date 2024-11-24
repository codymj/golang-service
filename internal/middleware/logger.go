package middleware

import (
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

// Middleware for logging request information.
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Grab logger for this request.
		logger := zerolog.Ctx(r.Context())

		// Note current time to measure request duration.
		start := time.Now()
		logger.Info().Msg("init")

		// Do work.
		next.ServeHTTP(w, r)

		// Log duration.
		logger.Info().
			Dur("duration", time.Duration(time.Since(start).Microseconds())).
			Msg("done")
	})
}
