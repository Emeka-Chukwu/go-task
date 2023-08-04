package auths

import (
	"context"
	domain "go-task/domain/auths/request"
	resp "go-task/domain/auths/response"
	"go-task/util"
	"time"
)

const dbTimeout = time.Second * 5

// register implements auths.Authentication.
func (auth *authentication) Register(data domain.RegisterModel) (*resp.RegisterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	hashedPassword, err := util.HashPassword(data.PasswordHash)
	if err != nil {
		return nil, err
	}
	stmt := `insert into users (username, email, password_hash)
		values ($1, $2, $3) returning id, username, email, created_at, updated_at`
	var user resp.RegisterResponse
	err = auth.DB.QueryRowContext(ctx, stmt,
		data.Username,
		data.Email,
		hashedPassword,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Email,
	)
	return &user, err
}
