package graph

import (
	"fmt"
	mockdb "go-task/internal/auths/usecase/mock"
	mockdbLabel "go-task/internal/labels/usecase/mock"
	mockdbTask "go-task/internal/tasks/usecase/mock"
	"go-task/token"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	// authorizationPayloadKey = "authorization_payload"
)

type UsecaseTest struct {
	Auth  *mockdb.MockAuthusecase
	Task  *mockdbTask.MockTaskusecase
	Label *mockdbLabel.MockLabelusecase
}

func newTestUsecase(t *testing.T) (Config, UsecaseTest) {

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
	return c, UsecaseTest{
		Auth:  auth,
		Label: label,
		Task:  task,
	}
}

func addAuthorization(
	t *testing.T,
	tokenMaker token.Maker,
	authorizationType string,
	username string,
	duration time.Duration,
) string {
	token, payload, err := tokenMaker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	return authorizationHeader
}

// client.AddHeader("Authorization")
