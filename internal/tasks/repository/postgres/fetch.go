package tasks

import (
	"context"
	resp "go-task/domain/task/response"
)

// Fetch implements Task.
func (t *task) FetchTask() ([]resp.TaskResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmp := `select id, title, description, status, priority, due_date, created_at, updated_at from tasks`
	tasks := make([]resp.TaskResponse, 0)
	rows, err := t.db.QueryContext(ctx, stmp)
	for rows.Next() {
		var model resp.TaskResponse
		rows.Scan(&model.ID, &model.Title, &model.Description,
			&model.Status, &model.Priority, &model.DueDate,
			&model.CreatedAt, &model.UpdatedAt)
		tasks = append(tasks, model)
	}
	return tasks, err
}
