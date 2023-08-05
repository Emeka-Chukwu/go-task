package domain

import (
	field "go-task/domain"
	"go-task/domain/auths"
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

type TaskModel struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	// Status      string `json:"status"`
	Priority string     `json:"priority"`
	DueDate  *time.Time `json:"due_date"`
}

var (
	priorities = []string{"low", "medium", "high"}
	statuses   = []string{"todo", "in-progress", "completes", "pending", "blocked", "rejected"}
)

func (req *TaskModel) Validate() (violations []*errdetails.BadRequest_FieldViolation) {
	if err := auths.ValidateNotEmptyString(req.Title); err != nil {
		violations = append(violations, field.FieldViolation("title", err))
	}
	if req.Description != nil {
		if err := auths.ValidateNotEmptyString(*req.Description); err != nil {
			violations = append(violations, field.FieldViolation("description", err))
		}
	}
	if err := auths.IsContained(req.Priority, priorities); err != nil {
		violations = append(violations, field.FieldViolation("priority", err))
	}

	if req.DueDate != nil {
		if err := auths.ValidateTimeStamp(*req.DueDate); err != nil {
			violations = append(violations, field.FieldViolation("due_date", err))
		}
	}
	return violations
}
func (req *UpdateTaskModel) Validate() (violations []*errdetails.BadRequest_FieldViolation) {
	if req.Title != nil {
		if err := auths.ValidateNotEmptyString(*req.Title); err != nil {
			violations = append(violations, field.FieldViolation("title", err))
		}
	}
	if req.Description != nil {
		if err := auths.ValidateNotEmptyString(*req.Description); err != nil {
			violations = append(violations, field.FieldViolation("description", err))
		}
	}
	if req.Priority != nil {
		if err := auths.IsContained(*req.Priority, priorities); err != nil {
			violations = append(violations, field.FieldViolation("priority", err))
		}
	}
	if req.Status != nil {
		if err := auths.IsContained(*req.Status, statuses); err != nil {
			violations = append(violations, field.FieldViolation("status", err))
		}
	}

	if req.DueDate != nil {
		if err := auths.ValidateTimeStamp(*req.DueDate); err != nil {
			violations = append(violations, field.FieldViolation("due_date", err))
		}
	}
	return violations
}

type UpdateTaskModel struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Status      *string    `json:"status"`
	Priority    *string    `json:"priority"`
	DueDate     *time.Time `json:"due_date"`
}
