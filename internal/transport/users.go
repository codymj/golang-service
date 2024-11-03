package transport

import (
	"time"

	"golang-service.codymj.io/internal/models"
)

// User data transfer object for client response.
type UserDTO struct {
	ID          int64     `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Location    *string   `json:"location,omitempty"`
	IsValidated bool      `json:"is_validated"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}

// Transforms models.User to UserDTO.
func ToUserDTO(user models.User) UserDTO {
	dto := UserDTO{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		IsValidated: user.IsValidated,
		CreatedAt:   time.UnixMilli(user.CreatedAt),
		ModifiedAt:  time.UnixMilli(user.ModifiedAt),
	}

	if user.Location.Valid {
		dto.Location = &user.Location.String
	}

	return dto
}
