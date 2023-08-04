package domain

import (
	field "go-task/domain"
	"go-task/domain/auths"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

type LabelModel struct {
	Name string `json:"name"`
}

func (req *LabelModel) ValidateLabel() (violations []*errdetails.BadRequest_FieldViolation) {
	if err := auths.ValidateNotEmptyString(req.Name); err != nil {
		violations = append(violations, field.FieldViolation("name", err))
	}
	return violations
}
