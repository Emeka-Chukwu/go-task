package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type LabelResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LabelTaskResponse struct {
	LabelResponse
	Tasks json.RawMessage `json:"tasks"`
}
