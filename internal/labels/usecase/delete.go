package label

import "github.com/google/uuid"

// Delete implements Labelusecase.
func (l *labelusecase) Delete(id uuid.UUID) error {
	return l.store.Delete(id)
}
