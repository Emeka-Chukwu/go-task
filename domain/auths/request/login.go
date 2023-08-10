package domain

import (
	field "go-task/domain"
	"go-task/domain/auths"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

type LoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *LoginModel) ValidateLogin() (violations []*errdetails.BadRequest_FieldViolation) {
	if err := auths.ValidatePassword(req.Password); err != nil {
		violations = append(violations, field.FieldViolation("password", err))
	}
	if err := auths.ValidateEmail(req.Email); err != nil {
		violations = append(violations, field.FieldViolation("email", err))
	}
	return violations
}
