package services

import (
	"context"
	"database/sql"

	"golang-service.codymj.io/internal/repos"
	"golang-service.codymj.io/internal/transport"
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
func (s *UsersService) List(ctx context.Context, username, email string) ([]transport.UserDTO, error) {
	// Retrieve users from database.
	users, err := s.repo.FindAll(ctx, repos.UsersRepoFindAllParams{
		Username: sql.NullString{String: username, Valid: username != ""},
		Email:    sql.NullString{String: email, Valid: email != ""},
	})
	if err != nil {
		return nil, err
	}

	// Transform user models to DTOs.
	response := make([]transport.UserDTO, 0)
	for _, user := range users {
		response = append(response, transport.ToUserDTO(user))
	}

	return response, nil
}
