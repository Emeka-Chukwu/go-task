package label

import resp "go-task/domain/label/response"

// List implements Labelusecase.
func (l *labelusecase) List() ([]resp.LabelResponse, error) {
	return l.store.List()
}
