package label

import (
	resp "go-task/domain/label/response"

	"github.com/google/uuid"
)

// ListByLabelID implements Labelusecase.
func (l *labelusecase) ListByLabelID(labelID uuid.UUID) (resp.LabelTaskResponse, error) {
	return l.store.ListByLabelID(labelID)
}
