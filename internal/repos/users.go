package repos

import (
	"context"
	"database/sql"
	"sync"

	sq "github.com/Masterminds/squirrel"
	"golang-service.codymj.io/internal/models"
)

// UsersRepo is the user repository.
type UsersRepo struct {
	_  sync.Mutex
	db *sql.DB
}

// NewUsersRepo returns a user repository.
func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

// UserRepoFindAllParams are params to filter with the FindAll query.
type UsersRepoFindAllParams struct {
	Username sql.NullString
	Email    sql.NullString
}

// FindAll finds all users by optional parameters.
func (r *UsersRepo) FindAll(ctx context.Context, params UsersRepoFindAllParams) ([]models.User, error) {
	qb := sq.Select(
		"id",
		"username",
		"email",
		"location",
		"is_validated",
		"created_at",
		"modified_at",
	).From(
		"users",
	)
	if params.Username.Valid {
		qb = qb.Where(sq.Eq{"username": params.Username.String})
	}
	if params.Email.Valid {
		qb = qb.Where(sq.Eq{"email": params.Email.String})
	}

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Location,
			&user.IsValidated,
			&user.CreatedAt,
			&user.ModifiedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
