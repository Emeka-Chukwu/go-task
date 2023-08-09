package tasks

import (
	"context"
	domain "go-task/domain/task/request"

	"go-task/token"
	"go-task/util"

	"github.com/google/uuid"

	repo "go-task/internal/tasks/repository/postgres"
)

type ResponseData struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

type Taskusecase interface {
	CreateTask(data domain.TaskModel) (ResponseData, error)
	UpdateTask(data domain.UpdateTaskModel, id uuid.UUID) (ResponseData, error)
	FetchTask() (ResponseData, error)
	FetchTaskByID(context context.Context, id uuid.UUID) (ResponseData, error)
	DeleteTask(id uuid.UUID) ResponseData
}

type taskusecase struct {
	store      repo.Task
	config     util.Config
	tokenMaker token.Maker
}

func NewTaskusecase(store repo.Task, config util.Config, tokenMaker token.Maker) Taskusecase {
	return &taskusecase{store: store, config: config, tokenMaker: tokenMaker}
}
