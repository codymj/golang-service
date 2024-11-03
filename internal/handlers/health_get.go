package handlers

import (
	"net/http"
)

// Application properties to display in the health response.
type HealthProperties struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Version   string `json:"version"`
}

// HTTP handler to show health status.
type HealthGetHandler struct {
	Properties HealthProperties
}

// Returns a HealthGetHandler.
func NewHealthGetHandler(props HealthProperties) *HealthGetHandler {
	return &HealthGetHandler{
		Properties: props,
	}
}

// Handles the HTTP request.
func (h *HealthGetHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Return response.
	headers := make(http.Header)

	err := writeJson(w, http.StatusOK, h.Properties, headers)
	if err != nil {
		serverErrorResp(w, r, err)
		return
	}
}
