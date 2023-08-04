package auths

import (
	domain "go-task/domain/auths/request"
	resp "go-task/domain/auths/response"
)

// login implements auths.Authentication.
func (auth *authentication) login(data domain.LoginModel) (resp.RegisterResponse, error) {
	panic("unimplemented")
}
