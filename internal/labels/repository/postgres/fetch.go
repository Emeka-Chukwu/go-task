package label

import (
	"context"
	resp "go-task/domain/label/response"

	"github.com/google/uuid"
)

// GetByID implements Label.
func (lab *label) GetByID(id uuid.UUID) (*resp.LabelResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmp := `selecte id, name, created_at, updated_at from labels where id=$1`
	var label resp.LabelResponse
	err := lab.DB.QueryRowContext(ctx, stmp, id).
		Scan(
			&label.ID,
			&label.Name,
			&label.CreatedAt,
			&label.UpdatedAt,
		)
	return &label, err
}
