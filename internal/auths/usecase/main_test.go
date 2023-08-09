package auths

import (
	repo "go-task/internal/auths/repository/postgres"
	"go-task/token"
	"go-task/util"
	"testing"
	"time"
)

func newTestUsecase(t *testing.T, store repo.Authentication) Authusecase {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil
	}
	usecase := NewAuthusecase(store, config, tokenMaker)
	return usecase
}
