package graph

import (
	resp "go-task/domain/auths/response"
	"go-task/graph/model"
)

func responseUser(user resp.LoginResponse) *model.LoginResponse {
	return &model.LoginResponse{
		User: &model.User{
			ID:        user.ID.String(),
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
		Token:     user.Token,
		ExpiredAt: user.ExpiredAt.String(),
	}
}

func responseUserData(user resp.RegisterResponse) *model.User {
	return &model.User{
		ID:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}
