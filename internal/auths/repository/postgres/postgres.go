package auths

import (
	"database/sql"

	domain "go-task/domain/auths/request"
	resp "go-task/domain/auths/response"
	"go-task/token"

	"github.com/google/uuid"
)

type authentication struct {
	DB         *sql.DB
	tokenMaker token.Maker
}

func NewAuthentication(db *sql.DB, tokenMaker token.Maker) Authentication {
	return &authentication{DB: db, tokenMaker: tokenMaker}

}

type Authentication interface {
	// googleSignIn(data any)
	// faceboolSignIn(data any)
	// appleSignIn(data any)
	// twoFactorAuth(data any)
	Login(data domain.LoginModel) (*resp.LoginResponse, error)
	Register(data domain.RegisterModel) (*resp.RegisterResponse, error)
	GetUserByID(id uuid.UUID) (*resp.RegisterResponse, error)
	GetUserByEmail(email string) (*resp.RegisterResponse, error)
}
