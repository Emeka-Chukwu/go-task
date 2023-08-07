package auths

import (
	resp "go-task/domain/auths/response"

	"github.com/google/uuid"
)

// GetUserByEmail implements Authusecase.
func (t *authusecase) GetUserByEmail(email string) (resp.RegisterResponse, error) {
	return t.store.GetUserByEmail(email)
}

// GetUserByID implements Authusecase.
func (t *authusecase) GetUserByID(id uuid.UUID) (resp.RegisterResponse, error) {
	return t.store.GetUserByID(id)
}
