package tasks

import (
	"context"
	domain "go-task/domain/task/request"
	resp "go-task/domain/task/response"

	"github.com/google/uuid"
)

// UpdateTask implements Task.
func (t *task) UpdateTask(data domain.UpdateTaskModel, id uuid.UUID) (resp.TaskResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `update tasks set title=$1, description=$2, status=$3, priority=$4, due_date=$5  where id=$6 returning id, 
	title, description, status, priority, due_date, created_at, updated_at`
	var model resp.TaskResponse
	err := t.db.QueryRowContext(ctx, stmt, data.Title, data.Description, &data.Status, &data.Priority, &data.DueDate, id).
		Scan(
			&model.ID,
			&model.Title,
			&model.Description,
			&model.Status,
			&model.Priority,
			&model.DueDate,
			&model.CreatedAt,
			&model.UpdatedAt,
		)
	return model, err
}
