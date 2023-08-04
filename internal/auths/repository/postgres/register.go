package auths

import (
	domain "go-task/domain/auths/request"
	resp "go-task/domain/auths/response"
)

// register implements auths.Authentication.
func (auth *authentication) register(data domain.RegisterModel) (resp.RegisterResponse, error) {
	panic("unimplemented")
}
