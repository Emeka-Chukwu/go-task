package label

import (
	"context"
	domain "go-task/domain/label/request"
	resp "go-task/domain/label/response"
)

// Create implements Label.
func (lab *label) Create(data domain.LabelModel) (resp.LabelResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmp := `insert into labels (name) values ($1) returning id, name, created_at, updated_at`
	var label resp.LabelResponse
	err := lab.DB.QueryRowContext(ctx, stmp, data.Name).
		Scan(
			&label.ID,
			&label.Name,
			&label.CreatedAt,
			&label.UpdatedAt,
		)
	return label, err
}
