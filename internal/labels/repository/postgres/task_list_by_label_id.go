package label

import (
	"context"
	resp "go-task/domain/label/response"

	"github.com/google/uuid"
)

// ListByLabelID implements Label.
func (l *label) ListByLabelID(labelID uuid.UUID) (resp.LabelTaskResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmp := `  SELECT l.*, json_agg(t.*) AS tasks
	FROM labels l
	JOIN task_labels tl ON tl.label_id = l.id
	 left JOIN tasks t ON t.id = tl.task_id
     WHERE tl.label_id = $1
	GROUP BY l.id`
	var result resp.LabelTaskResponse
	err := l.DB.QueryRowContext(ctx, stmp, labelID).
		Scan(
			&result.ID,
			&result.Name,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.Tasks,
		)
	return result, err
}
