package domain

import (
	"time"

	"github.com/google/uuid"
)

type RegisterResponse struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type LoginResponse struct {
	RegisterResponse
	Token string `json:"token"`
}
