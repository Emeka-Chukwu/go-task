package tasks

import (
	repo "go-task/internal/tasks/repository/postgres"
	"go-task/token"
	"go-task/util"
	"testing"
	"time"
)

func newTestUsecase(t *testing.T, store repo.Task) Taskusecase {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil
	}
	usecase := NewTaskusecase(store, config, tokenMaker)
	return usecase
}
