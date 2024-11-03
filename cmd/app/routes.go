package app

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang-service.codymj.io/internal/handlers"
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
	healthProps := handlers.HealthProperties{
		Namespace: a.cfg.App.Namespace,
		Name:      a.cfg.App.Name,
		Version:   a.cfg.App.Version,
	}
	healthGetHandler := handlers.NewHealthGetHandler(healthProps)
	router.HandlerFunc(http.MethodGet, "/v1/health", healthGetHandler.Handle)

	usersGetHandler := handlers.NewUsersGetHandler(usersService)
	router.HandlerFunc(http.MethodGet, "/v1/users", usersGetHandler.Handle)

	return router
}
