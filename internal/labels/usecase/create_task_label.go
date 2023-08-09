package label

import domain "go-task/domain/label/request"

// CreateTaskLabel implements Labelusecase.
func (l *labelusecase) CreateTaskLabel(data domain.LabelTaskModel) error {
	return l.store.CreateTaskLabel(data)
}
