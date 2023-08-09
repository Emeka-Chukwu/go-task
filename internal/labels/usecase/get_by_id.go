package label

import (
	resp "go-task/domain/label/response"

	"github.com/google/uuid"
)

// GetByID implements Labelusecase.
func (l *labelusecase) GetByID(id uuid.UUID) (resp.LabelResponse, error) {
	return l.store.GetByID(id)
}
