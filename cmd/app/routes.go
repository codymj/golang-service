package app

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang-service.codymj.io/internal/handlers"
	"golang-service.codymj.io/internal/middleware"
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
	router.HandlerFunc(
		http.MethodGet,
		"/v1/health",
		makeChain(
			healthGetHandler.Handle,
			middleware.Logger,
			middleware.Tracer,
		),
	)

	usersGetHandler := handlers.NewUsersGetHandler(usersService)
	router.HandlerFunc(
		http.MethodGet,
		"/v1/users",
		makeChain(
			usersGetHandler.Handle,
			middleware.Logger,
			middleware.Tracer,
		),
	)

	return router
}

// Creates middleware chain around a handler.
// The first middleware in the chain will be executed last. For example, the
// order of execution of the middlewares wrapped by makeChain(h, mw1, mw2, mw3)
// will be mw3, mw2, mw1.
func makeChain(h http.HandlerFunc, mws ...middleware.Middleware) http.HandlerFunc {
	for _, mw := range mws {
		h = mw(h)
	}

	return h
}
