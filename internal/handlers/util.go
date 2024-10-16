package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

// Writes JSON response back to client.
func writeJson(w http.ResponseWriter, status int, data any, headers http.Header) error {
	// encode data
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	js = append(js, '\n')

	// add headers
	for k, v := range headers {
		w.Header()[k] = v
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(js)

	return nil
}

// Wraps responses.
type envelope map[string]any

// Encodes errors to JSON for sending to client.
func errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	e := envelope{
		"error": message,
	}
	err := writeJson(w, status, e, nil)
	if err != nil {
		log.Error().Msgf("%v %v", err, map[string]string{
			"requestMethod": r.Method,
			"requestUrl":    r.URL.String(),
		})
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Sends a 500 response to client.
func serverErrorResp(w http.ResponseWriter, r *http.Request, err error) {
	log.Error().Msgf("%v %v", err, map[string]string{
		"requestMethod": r.Method,
		"requestUrl":    r.URL.String(),
	})

	msg := "internal server error"
	errorResponse(w, r, http.StatusInternalServerError, msg)
}
