package label

import (
	"context"
	resp "go-task/domain/label/response"
)

// List implements Label.
func (lab *label) List() ([]resp.LabelResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmp := `select id, name, created_at, updated_at from labels`
	labels := make([]resp.LabelResponse, 0)
	rows, err := lab.DB.QueryContext(ctx, stmp)
	for rows.Next() {
		var model resp.LabelResponse
		rows.Scan(&model.ID, &model.Name, &model.CreatedAt, &model.UpdatedAt)
		labels = append(labels, model)
	}
	return labels, err
}
