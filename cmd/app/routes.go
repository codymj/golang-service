package app

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang-service.codymj.io/internal/handlers/users"
	"golang-service.codymj.io/internal/repos"
	"golang-service.codymj.io/internal/services"
)

// Creates and initializes the router.
func (a *application) routes(mariadb *sql.DB) http.Handler {
	// Initialize router.
	router := httprouter.New()

	// Initialize repositories.
	usersRepo := repos.NewUsersRepo(mariadb)

	// Initialize services.
	usersService := services.NewUsersService(usersRepo)

	// Initialize routes.
	usersListHandler := users.NewUsersListHandler(usersService)
	router.HandlerFunc(http.MethodGet, "/v1/users", usersListHandler.Handle)

	return router
}
