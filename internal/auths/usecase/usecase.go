package auths

import (
	domain "go-task/domain/auths/request"
	resp "go-task/domain/auths/response"
	repo "go-task/internal/auths/repository/postgres"
	"go-task/token"
	"go-task/util"

	"github.com/google/uuid"
)

type Authusecase interface {
	LoginUser(req domain.LoginModel) (resp.LoginResponse, error)
	Register(data domain.RegisterModel) (resp.LoginResponse, error)
	GetUserByID(id uuid.UUID) (resp.RegisterResponse, error)
	GetUserByEmail(email string) (resp.RegisterResponse, error)
}

type authusecase struct {
	store      repo.Authentication
	config     util.Config
	tokenMaker token.Maker
}

func NewAuthusecase(store repo.Authentication, config util.Config, tokenMaker token.Maker) Authusecase {
	return &authusecase{store: store, config: config, tokenMaker: tokenMaker}
}
