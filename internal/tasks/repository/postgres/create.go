package tasks

import (
	domain "go-task/domain/task/request"
	resp "go-task/domain/task/response"
)

// CreateTask implements Task.
func (*task) CreateTask(data domain.TaskModel) (*resp.TaskResponse, error) {
	panic("unimplemented")
}
