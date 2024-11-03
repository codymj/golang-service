package handlers

import (
	"net/http"

	"golang-service.codymj.io/internal/services"
)

// HTTP handler to list all users.
type UsersGetHandler struct {
	usersService *services.UsersService
}

// Returns a UsersGetHandler.
func NewUsersGetHandler(usersService *services.UsersService) *UsersGetHandler {
	return &UsersGetHandler{
		usersService: usersService,
	}
}

// Handles the HTTP request.
func (h *UsersGetHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Get query parameter map and populate query parameters, if any.
	queryMap := r.URL.Query()

	username := queryMap.Get("username")
	email := queryMap.Get("email")

	// Call service.
	users, err := h.usersService.List(r.Context(), username, email)
	if err != nil {
		serverErrorResp(w, r, err)
		return
	}

	// Return response.
	headers := make(http.Header)

	if len(users) == 0 {
		err = writeJson(w, http.StatusNoContent, nil, headers)
		if err != nil {
			serverErrorResp(w, r, err)
			return
		}
	} else {
		err = writeJson(w, http.StatusOK, users, headers)
		if err != nil {
			serverErrorResp(w, r, err)
			return
		}
	}
}
