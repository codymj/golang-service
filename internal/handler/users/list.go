package users

import (
	"net/http"

	"golang-service.codymj.io/internal/handler"
	"golang-service.codymj.io/internal/service"
)

// UsersListHandler is the HTTP handler to list all users.
type UsersListHandler struct {
	userSvc *service.UserService
}

// NewUsersListHandler returns a new UsersListHandler.
func NewUsersListHandler(userSvc *service.UserService) *UsersListHandler {
	return &UsersListHandler{
		userSvc: userSvc,
	}
}

// Handle handles the HTTP request.
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
	err = handler.WriteJson(w, http.StatusOK, users, headers)
	if err != nil {
		// handle
	}
}
