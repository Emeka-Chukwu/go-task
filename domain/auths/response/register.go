package domain

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/satori-uuid"
)

type RegisterResponse struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
