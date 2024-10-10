package services

import (
	"context"
	"database/sql"

	"golang-service.codymj.io/configs"
	"golang-service.codymj.io/internal/models"
	"golang-service.codymj.io/internal/repos"
)

// UsersService is the service to manage users.
type UsersService struct {
	cfg  *configs.Config
	repo *repos.UsersRepo
}

// NewUserService returns a new user service.
func NewUsersService(cfg *configs.Config, repo *repos.UsersRepo) *UsersService {
	return &UsersService{
		cfg:  cfg,
		repo: repo,
	}
}

// List returns a list of all users, filterable by optional parameters.
func (s *UsersService) List(ctx context.Context, username, email string) ([]models.User, error) {
	return s.repo.FindAll(ctx, repos.UsersRepoFindAllParams{
		Username: sql.NullString{String: username, Valid: username != ""},
		Email:    sql.NullString{String: email, Valid: email != ""},
	})
}
