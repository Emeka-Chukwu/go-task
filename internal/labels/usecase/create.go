package label

import (
	domain "go-task/domain/label/request"
	resp "go-task/domain/label/response"
)

// Create implements Labelusecase.
func (l *labelusecase) Create(data domain.LabelModel) (resp.LabelResponse, error) {
	value, err := l.store.Create(data)
	return value, err
}
