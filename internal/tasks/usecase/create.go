package tasks

import (
	"errors"
	domain "go-task/domain/task/request"
)

// CreateTask implements Taskusecase.
func (t *taskusecase) CreateTask(data domain.TaskModel) (resp ResponseData, err error) {
	errdetails := data.Validate()
	if errdetails != nil {
		resp.Error = errdetails
		resp.Message = "validation error"
		err = errors.New("validation error")
		return
	}
	task, err := t.store.CreateTask(data)
	if err != nil {
		resp.Error = err
		resp.Message = "internal error"
		return
	}
	resp.Data = task
	resp.Message = "success"
	return
}
