package graph

import (
	"context"
	"go-task/graph/model"
	useAuth "go-task/internal/auths/usecase"
	useLabel "go-task/internal/labels/usecase"
	useTask "go-task/internal/tasks/usecase"
)

type TestGraphQlMethods interface {
	CreateUser(ctx context.Context, input model.NewUser) (*model.LoginResponse, error)
	LoginUser(ctx context.Context, input model.LoginUser) (*model.LoginResponse, error)
}

func NewTestGraphQlMethods(auth useAuth.Authusecase, label useLabel.Labelusecase, task useTask.Taskusecase) TestGraphQlMethods {
	return &mutationResolver{&Resolver{Auth: auth, Label: label, Task: task}}
}

type TestGraphQlMethodsQuery interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
}
