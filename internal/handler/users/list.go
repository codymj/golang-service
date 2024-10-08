package users

import (
	"net/http"

	"golang-service.codymj.io/internal/handler"
	"golang-service.codymj.io/internal/service"
)

// ListUsersHandler is the HTTP handler to list all users.
type ListUsersHandler struct {
	service *service.UserService
}

// NewListUsersHandler returns a new ListUsersHandler.
func NewListUsersHandler(service *service.UserService) *ListUsersHandler {
	return &ListUsersHandler{
		service: service,
	}
}

// Handle handles the HTTP request.
func (h *ListUsersHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Get query string and populate query parameters, if any.
	qs := r.URL.Query()

	username := qs.Get("username")
	email := qs.Get("email")

	users, err := h.service.List(r.Context(), username, email)
	if err != nil {
		// handle
	}

	headers := make(http.Header)
	err = handler.WriteJson(w, http.StatusOK, users, headers)
	if err != nil {
		// handle
	}
}
