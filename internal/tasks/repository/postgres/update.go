package tasks

import (
	domain "go-task/domain/task/request"
	resp "go-task/domain/task/response"

	"github.com/google/uuid"
)

// UpdateTask implements Task.
func (*task) UpdateTask(data domain.TaskModel, id uuid.UUID) (*resp.TaskResponse, error) {
	panic("unimplemented")
}
