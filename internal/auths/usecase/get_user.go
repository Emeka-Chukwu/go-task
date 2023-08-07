package auths

import (
	resp "go-task/domain/auths/response"

	"github.com/google/uuid"
)

// GetUserByEmail implements Authusecase.
func (t *authusecase) GetUserByEmail(email string) (resp.RegisterResponse, error) {
	panic("unimplemented")
}

// GetUserByID implements Authusecase.
func (*authusecase) GetUserByID(id uuid.UUID) (resp.RegisterResponse, error) {
	panic("unimplemented")
}
