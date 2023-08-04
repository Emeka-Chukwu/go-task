package auths

import (
	"database/sql"

	domain "go-task/domain/auths/request"
	resp "go-task/domain/auths/response"
)

type authentication struct {
	DB *sql.DB
}

func NewAuthentication(db *sql.DB) Authentication {
	return &authentication{DB: db}

}

type Authentication interface {
	// googleSignIn(data any)
	// faceboolSignIn(data any)
	// appleSignIn(data any)
	// twoFactorAuth(data any)
	register(data domain.RegisterModel) (resp.RegisterResponse, error)
	login(data domain.LoginModel) (resp.RegisterResponse, error)
}
