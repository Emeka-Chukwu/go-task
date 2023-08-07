package auths

import (
	domain "go-task/domain/auths/request"
	resp "go-task/domain/auths/response"
	"go-task/util"
)

// Register implements Authusecase.
func (t *authusecase) Register(data domain.RegisterModel) (resp.LoginResponse, error) {
	HashPassword, err := util.HashPassword(data.PasswordHash)
	if err != nil {
		return resp.LoginResponse{}, err
	}
	data.PasswordHash = HashPassword
	user, err := t.store.Register(data)
	if err != nil {
		return resp.LoginResponse{}, err
	}
	token, payload, err := t.tokenMaker.CreateToken(user.ID.String(), t.config.AccessTokenDuration)
	if err != nil {
		return resp.LoginResponse{}, err
	}
	return resp.LoginResponse{
		Token:            token,
		RegisterResponse: user,
		ExpiredAt:        payload.ExpiredAt,
	}, err
}
