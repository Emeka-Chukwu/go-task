package tasks

import (
	"context"
	resp "go-task/domain/task/response"

	"github.com/google/uuid"
)

// FetchTaskByID implements Task.

var stmt = `select id, title, description, status, priority, due_date, created_at, updated_at from tasks where id=$1 `

func (t *task) FetchTaskByID(context context.Context, id uuid.UUID) (resp.TaskResponse, error) {
	var task resp.TaskResponse
	err := t.db.QueryRowContext(context, stmt, id).
		Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.Priority,
			&task.DueDate,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
	return task, err
}
