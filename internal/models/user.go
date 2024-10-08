package models

import "database/sql"

// User data model.
type User struct {
	ID          int64          `json:"id"`
	Username    string         `json:"username"`
	Email       string         `json:"email"`
	Password    string         `json:"password,omitempty"`
	Location    sql.NullString `json:"location,omitempty"`
	IsValidated bool           `json:"is_validated"`
	CreatedAt   int64          `json:"created_at"`
	ModifiedAt  int64          `json:"modified_at"`
}
