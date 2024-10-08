package service

import (
	"context"
	"database/sql"

	"golang-service.codymj.io/config"
	"golang-service.codymj.io/internal/model"
	"golang-service.codymj.io/internal/repo"
)

// UserService is the service to manage users.
type UserService struct {
	cfg  *config.Config
	repo *repo.UserRepo
}

// NewUserService returns a new user service.
func NewUserService(cfg *config.Config, repo *repo.UserRepo) *UserService {
	return &UserService{
		cfg:  cfg,
		repo: repo,
	}
}

// List returns a list of all users, filterable by optional parameters.
func (s *UserService) List(ctx context.Context, username, email string) ([]model.User, error) {
	return s.repo.FindAll(ctx, repo.UserRepoFindAllParams{
		Username: sql.NullString{String: username, Valid: username != ""},
		Email:    sql.NullString{String: email, Valid: email != ""},
	})
}
