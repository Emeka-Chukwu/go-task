package tasks

import (
	"context"
	domain "go-task/domain/task/request"
	resp "go-task/domain/task/response"
)

// CreateTask implements Task.
func (t *task) CreateTask(data domain.TaskModel) (resp.TaskResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `insert into tasks ( title , description, status, priority , due_date, user_id ) values ($1, $2, $3, $4, $5, $6) returning id, 
	title, description, status, priority, due_date, created_at, updated_at`
	var response resp.TaskResponse
	err := t.db.QueryRowContext(ctx, stmt,
		data.Title,
		data.Description,
		"todo",
		data.Priority,
		data.DueDate,
		data.UserID,
	).Scan(
		&response.ID,
		&response.Title,
		&response.Description,
		&response.Status,
		&response.Priority,
		&response.DueDate,
		&response.CreatedAt,
		&response.UpdatedAt,
	)
	return response, err
}
