package services

import (
	"context"
	"database/sql"

	"golang-service.codymj.io/internal/models"
	"golang-service.codymj.io/internal/repos"
)

// Service to manage users.
type UsersService struct {
	repo *repos.UsersRepo
}

// Returns a new user service.
func NewUsersService(repo *repos.UsersRepo) *UsersService {
	return &UsersService{
		repo: repo,
	}
}

// Returns a list of all users, filterable by optional parameters.
func (s *UsersService) List(ctx context.Context, username, email string) ([]models.User, error) {
	return s.repo.FindAll(ctx, repos.UsersRepoFindAllParams{
		Username: sql.NullString{String: username, Valid: username != ""},
		Email:    sql.NullString{String: email, Valid: email != ""},
	})
}
