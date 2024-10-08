package handler

import (
	"encoding/json"
	"net/http"
)

// WriteJson writes JSON response back to client.
func WriteJson(w http.ResponseWriter, status int, data any, headers http.Header) error {
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
