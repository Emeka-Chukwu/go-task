package label

import (
	"context"
	domain "go-task/domain/label/request"
)

// CreateTaskLabel implements Label.
func (l *label) CreateTaskLabel(data domain.LabelTaskModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmp := `insert into task_labels (task_id, label_id) values ($1, $2) returning id,task_id , label_id, created_at`
	err := l.DB.QueryRowContext(ctx, stmp, data.TaskID, data.LabelID)
	return err.Err()
}
