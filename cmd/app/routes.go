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
	// Initialize router.
	router := httprouter.New()

	// Initialize repositories.
	userRepo := repo.NewUserRepo(db)

	// Initialize services.
	userService := service.NewUserService(a.cfg, userRepo)

	// Initialize routes.
	usersListHandler := users.NewUsersListHandler(userService)
	router.HandlerFunc(http.MethodGet, "/v1/users", usersListHandler.Handle)

	return router
}
