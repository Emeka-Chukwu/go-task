package auths

import (
	domain "go-task/domain/auths/request"
	resp "go-task/domain/auths/response"
	"go-task/util"
)

// LoginUser implements Authusecase.
func (t *authusecase) LoginUser(req domain.LoginModel) (resp.LoginResponse, error) {
	user, err := t.store.GetUserByEmail(req.Email)
	if err != nil {
		return resp.LoginResponse{}, err
	}
	if err := util.CheckPassword(req.Password, user.PasswordHash); err != nil {
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
