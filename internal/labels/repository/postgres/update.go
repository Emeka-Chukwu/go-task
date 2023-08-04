package label

import (
	"context"
	domain "go-task/domain/label/request"
	resp "go-task/domain/label/response"

	"github.com/google/uuid"
)

// Update implements Label.
func (lab *label) Update(id uuid.UUID, data domain.LabelModel) (*resp.LabelResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `update labels set name:$1 where id=$2 returning id, name, created_at, updated_at`
	var model resp.LabelResponse
	err := lab.DB.QueryRowContext(ctx, stmt, data.Name, id).
		Scan(
			&model.ID,
			&model.Name,
			&model.CreatedAt,
			&model.UpdatedAt,
		)
	return &model, err
}
