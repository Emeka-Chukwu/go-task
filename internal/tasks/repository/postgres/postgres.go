package tasks

import (
	"database/sql"
	domain "go-task/domain/task/request"
	resp "go-task/domain/task/response"
	"go-task/token"

	"github.com/google/uuid"
)

type Task interface {
	CreateTask(data domain.TaskModel) (*resp.TaskResponse, error)
	UpdateTask(data domain.TaskModel, id uuid.UUID) (*resp.TaskResponse, error)
	FetchTask() ([]*resp.TaskResponse, error)
	DeleteTask(id uuid.UUID) error
}

type task struct {
	DB    *sql.DB
	token token.Maker
}

func NewTask(db *sql.DB, token token.Maker) Task {
	return &task{DB: db, token: token}
}
