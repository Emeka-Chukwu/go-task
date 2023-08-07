package auths

import (
	"context"
	"fmt"
	domain "go-task/domain/auths/request"
	resp "go-task/domain/auths/response"
	"time"
)

const dbTimeout = time.Second * 5

// register implements auths.Authentication.
func (auth *authentication) Register(data domain.RegisterModel) (resp.RegisterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `insert into users (username, email, password_hash)
		values ($1, $2, $3) returning id, username, email, password_hash, created_at, updated_at`
	var user resp.RegisterResponse
	err := auth.DB.QueryRowContext(ctx, stmt,
		data.Username,
		data.Email,
		data.PasswordHash,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	fmt.Println(user.PasswordHash)
	return user, err
}
