package auths

import (
	domain "go-task/domain/auths/request"
	resp "go-task/domain/auths/response"
)

// Register implements Authusecase.
func (*authusecase) Register(data domain.RegisterModel) (resp.RegisterResponse, error) {
	panic("unimplemented")
}
