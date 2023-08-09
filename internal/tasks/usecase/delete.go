package tasks

import "github.com/google/uuid"

// DeleteTask implements Taskusecase.
func (t *taskusecase) DeleteTask(id uuid.UUID) ResponseData {
	err := t.store.DeleteTask(id)
	if err != nil {
		return ResponseData{Message: "internal err", Error: err}
	}
	return ResponseData{Message: "successfully", Data: ""}
}
