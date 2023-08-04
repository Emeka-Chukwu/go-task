package domain

import (
	"go-task/domain/auths"

	field "go-task/domain"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

type RegisterModel struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func (req *RegisterModel) ValidateRegister() (violations []*errdetails.BadRequest_FieldViolation) {
	if err := auths.ValidatePassword(req.PasswordHash); err != nil {
		violations = append(violations, field.FieldViolation("password", err))
	}
	if err := auths.ValidateEmail(req.Email); err != nil {
		violations = append(violations, field.FieldViolation("email", err))
	}
	if err := auths.ValidateNotEmptyString(req.Username); err != nil {
		violations = append(violations, field.FieldViolation("username", err))
	}
	return violations
}
