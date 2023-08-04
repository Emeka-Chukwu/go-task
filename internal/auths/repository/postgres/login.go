package auths

import (
	"context"
	domain "go-task/domain/auths/request"
	resp "go-task/domain/auths/response"
	"go-task/util"
	"time"
)

func (auth *authentication) Login(data domain.LoginModel) (*resp.LoginResponse, error) {
	_, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	user, err := auth.GetUserByEmail(data.Email)
	if err != nil {
		return nil, err
	}
	if err := util.CheckPassword(data.Password, user.PasswordHash); err != nil {
		return nil, err
	}
	token, _, err := auth.tokenMaker.CreateToken(user.ID.UUID.String(), time.Hour)
	if err != nil {
		return nil, err
	}
	if err := util.CheckPassword(data.Password, user.PasswordHash); err != nil {
		return nil, err
	}
	value := resp.LoginResponse{Token: token, RegisterResponse: *user}
	return &value, nil
}
