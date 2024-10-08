package app

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang-service.codymj.io/internal/handler/users"
	"golang-service.codymj.io/internal/repo"
	"golang-service.codymj.io/internal/service"
)

func (a *application) routes(db *sql.DB) http.Handler {
	// Intialize router.
	router := httprouter.New()

	// Setup /v1/users routes.
	userRepo := repo.NewUserRepo(db)
	userService := service.NewUserService(a.cfg, userRepo)

	listUsersHandler := users.NewListUsersHandler(userService)
	router.HandlerFunc(http.MethodGet, "/v1/users", listUsersHandler.Handle)

	return router
}
