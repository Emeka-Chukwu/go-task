package domain

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/satori-uuid"
)

type LabelResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
