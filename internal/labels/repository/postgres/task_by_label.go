package label

import (
	"context"
	resp "go-task/domain/label/response"
)

// ListByLabel implements Label.
func (l *label) ListByLabel() ([]resp.LabelTaskResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmp := ` SELECT l.id, l.name, l.created_at, l.updated_at, json_agg(t.*) AS tasks
	FROM labels l
	 JOIN task_labels tl ON tl.label_id = l.id
	 JOIN tasks t ON t.id = tl.task_id
	GROUP BY l.id`
	labels := make([]resp.LabelTaskResponse, 0)
	rows, err := l.DB.QueryContext(ctx, stmp)
	for rows.Next() {
		var result resp.LabelTaskResponse
		rows.Scan(&result.ID, &result.Name, &result.CreatedAt, &result.UpdatedAt, &result.Tasks)
		labels = append(labels, result)
	}
	return labels, err
}
