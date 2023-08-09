package label

import (
	repo "go-task/internal/labels/repository/postgres"
	"go-task/token"
	"go-task/util"
	"testing"
	"time"
)

func newTestUsecase(t *testing.T, store repo.Label) Labelusecase {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil
	}
	usecase := NewLabelusecase(store, config, tokenMaker)
	return usecase
}
