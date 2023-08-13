package graph

import (
	"fmt"
	resp "go-task/domain/auths/response"
	respTask "go-task/domain/task/response"
	"go-task/graph/model"
)

func responseUser(user resp.LoginResponse) *model.LoginResponse {
	return &model.LoginResponse{
		User: &model.User{
			ID:        user.ID.String(),
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
		Token:     user.Token,
		ExpiredAt: user.ExpiredAt.String(),
	}
}

func responseUserData(user resp.RegisterResponse) *model.User {
	return &model.User{
		ID:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}

func responseTaskData(task respTask.TaskResponse) *model.TaskResponse {
	date := fmt.Sprintf("&v", task.DueDate)

	return &model.TaskResponse{
		Message: "Task fetched successfully",
		Data: &model.Task{
			ID:          task.ID.String(),
			Title:       task.Title,
			Description: *task.Description,
			DueDate:     &date,
			Priority:    task.Priority,
			CreatedAt:   task.CreatedAt.GoString(),
			UpdatedAt:   task.UpdatedAt.GoString(),
			Status:      task.Status,
		},
	}
}

func responseListTaskData(tasks []respTask.TaskResponse) *model.TaskListResponse {
	tasksResp := make([]*model.Task, 0)
	for _, task := range tasks {
		date := fmt.Sprintf("&v", task.DueDate)
		taskData := &model.Task{
			ID:          task.ID.String(),
			Title:       task.Title,
			Description: *task.Description,
			DueDate:     &date,
			Priority:    task.Priority,
			CreatedAt:   task.CreatedAt.GoString(),
			UpdatedAt:   task.UpdatedAt.GoString(),
			Status:      task.Status,
		}
		tasksResp = append(tasksResp, taskData)
	}
	return &model.TaskListResponse{
		Message: "Task fetched successfully",
		Data:    tasksResp,
	}
}
