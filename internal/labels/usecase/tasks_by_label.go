package label

import resp "go-task/domain/label/response"

// ListByLabel implements Labelusecase.
func (l *labelusecase) ListByLabel() ([]resp.LabelTaskResponse, error) {
	return l.store.ListByLabel()
}
