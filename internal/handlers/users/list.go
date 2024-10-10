package users

import (
	"net/http"

	"golang-service.codymj.io/internal/handlers"
	"golang-service.codymj.io/internal/services"
)

// The HTTP handler to list all users.
type UsersListHandler struct {
	userSvc *services.UsersService
}

// Returns a UsersListHandler.
func NewUsersListHandler(userSvc *services.UsersService) *UsersListHandler {
	return &UsersListHandler{
		userSvc: userSvc,
	}
}

// Handles the HTTP request.
func (h *UsersListHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Get query parameter map and populate query parameters, if any.
	queryMap := r.URL.Query()

	username := queryMap.Get("username")
	email := queryMap.Get("email")

	// Call service.
	users, err := h.userSvc.List(r.Context(), username, email)
	if err != nil {
		// handle
	}

	// Return response.
	headers := make(http.Header)
	err = handlers.WriteJson(w, http.StatusOK, users, headers)
	if err != nil {
		// handle
	}
}
