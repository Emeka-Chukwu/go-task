package graph

import (
	mockdb "go-task/internal/auths/usecase/mock"
	mockdbLabel "go-task/internal/labels/usecase/mock"
	mockdbTask "go-task/internal/tasks/usecase/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func newTestUsecase(t *testing.T) TestGraphQlMethods {
	// config := util.Config{
	// 	TokenSymmetricKey:   util.RandomString(32),
	// 	AccessTokenDuration: time.Minute,
	// }
	// tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	// if err != nil {
	// 	return nil
	// }
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctrl2 := gomock.NewController(t)
	defer ctrl.Finish()
	ctrl3 := gomock.NewController(t)
	defer ctrl.Finish()
	auth := mockdb.NewMockAuthusecase(ctrl)
	label := mockdbLabel.NewMockLabelusecase(ctrl2)
	task := mockdbTask.NewMockTaskusecase(ctrl3)
	usecase := NewTestGraphQlMethods(auth, label, task)
	return usecase
}
