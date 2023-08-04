package auths

import (
	domain "go-task/domain/auths/request"
	resp "go-task/domain/auths/response"
)

type Authentication interface {
	// googleSignIn(data any)
	// faceboolSignIn(data any)
	// appleSignIn(data any)
	// twoFactorAuth(data any)
	register(data domain.RegisterModel) (resp.RegisterResponse, error)
}
