package auths

import (
	"context"
	resp "go-task/domain/auths/response"

	"github.com/google/uuid"
)

// GetUserByEmail implements Authentication.
func (auth *authentication) GetUserByEmail(email string) (*resp.RegisterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `select id, username, email, password_hash, created_at, updated_at where email=$1`
	var user resp.RegisterResponse
	err := auth.DB.QueryRowContext(ctx, stmt, email).
		Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.PasswordHash,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.Email,
		)
	return &user, err
}

// GetUserByID implements Authentication.
func (auth *authentication) GetUserByID(id uuid.UUID) (*resp.RegisterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `select id, username, email, password_hash, created_at, updated_at where id=$1`
	var user resp.RegisterResponse
	err := auth.DB.QueryRowContext(ctx, stmt, id).
		Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.PasswordHash,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.Email,
		)
	return &user, err
}
