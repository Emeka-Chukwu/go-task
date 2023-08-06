package domain

import (
	field "go-task/domain"
	"go-task/domain/auths"

	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

type LabelModel struct {
	Name   string  `json:"name"`
	TaskID *string `json:"task_id"`
}

func (req *LabelModel) ValidateLabel() (violations []*errdetails.BadRequest_FieldViolation) {
	if err := auths.ValidateNotEmptyString(req.Name); err != nil {
		violations = append(violations, field.FieldViolation("name", err))
	}
	if req.TaskID != nil {
		if err := auths.ValidateUUID(*req.TaskID); err != nil {
			violations = append(violations, field.FieldViolation("task_id", err))
		}
	}
	return violations
}

type LabelTaskModel struct {
	LabelID uuid.UUID `json:"label_id"`
	TaskID  uuid.UUID `json:"task_id"`
}
