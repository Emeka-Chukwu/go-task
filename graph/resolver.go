package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	useAuth "go-task/internal/auths/usecase"
	useLabel "go-task/internal/labels/usecase"
	useTask "go-task/internal/tasks/usecase"
)

type Resolver struct {
	Auth  useAuth.Authusecase
	Label useLabel.Labelusecase
	Task  useTask.Taskusecase
}
