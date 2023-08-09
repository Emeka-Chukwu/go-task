package tasks

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

// FetchTaskByID implements Taskusecase.
func (t *taskusecase) FetchTaskByID(context context.Context, id uuid.UUID) (ResponseData, error) {
	resp, err := t.store.FetchTaskByID(context, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ResponseData{Message: "Record not found"}, err
		}
		return ResponseData{Message: "Internal server err"}, err
	}
	return ResponseData{Message: "Data fetched successfully", Data: resp}, nil
}
