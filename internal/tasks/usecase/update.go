package tasks

import (
	"database/sql"
	"errors"
	domain "go-task/domain/task/request"

	"github.com/google/uuid"
)

// UpdateTask implements Taskusecase.
func (t *taskusecase) UpdateTask(data domain.UpdateTaskModel, id uuid.UUID) (ResponseData, error) {
	errdetails := data.Validate()
	if errdetails != nil {
		return ResponseData{Message: "Validation err", Error: errdetails}, errors.New("validation err")
	}
	updatedTask, err := t.store.UpdateTask(data, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ResponseData{Message: "Record not found"}, err
		}
		return ResponseData{Message: "Internal server error", Error: err}, err
	}
	return ResponseData{Message: "Record updated successfully", Data: updatedTask}, err
}
