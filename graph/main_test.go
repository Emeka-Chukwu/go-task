package graph

import (
	mockdb "go-task/internal/auths/usecase/mock"
	mockdbLabel "go-task/internal/labels/usecase/mock"
	mockdbTask "go-task/internal/tasks/usecase/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func newTestUsecase(t *testing.T) (Config, *mockdb.MockAuthusecase) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctrl2 := gomock.NewController(t)
	defer ctrl.Finish()
	ctrl3 := gomock.NewController(t)
	defer ctrl.Finish()
	auth := mockdb.NewMockAuthusecase(ctrl)
	label := mockdbLabel.NewMockLabelusecase(ctrl2)
	task := mockdbTask.NewMockTaskusecase(ctrl3)

	c := Config{Resolvers: &Resolver{
		Auth:  auth,
		Label: label,
		Task:  task,
	}}
	// c.Directives.Auth = directives.Auth
	return c, auth
}
