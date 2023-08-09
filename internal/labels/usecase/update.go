package label

import (
	domain "go-task/domain/label/request"
	resp "go-task/domain/label/response"

	"github.com/google/uuid"
)

// Update implements Labelusecase.
func (l *labelusecase) Update(id uuid.UUID, data domain.LabelModel) (resp.LabelResponse, error) {
	return l.store.Update(id, data)
}
