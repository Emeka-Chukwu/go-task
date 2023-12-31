package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"fmt"
	domain "go-task/domain/task/request"
	resp "go-task/domain/task/response"
	"go-task/graph/model"
	"go-task/middlewares"
	"time"

	"github.com/google/uuid"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.TaskResponse, error) {
	var dueDate time.Time
	var err error
	if input.DueDate != nil {
		const layout = "2006-01-02 15:04:05"
		date := input.DueDate
		dueDate, err = time.Parse(layout, *date)
		if err != nil {
			return nil, fmt.Errorf("invalid duedate")
		}
	}

	payload, _ := middlewares.GetCurrentUserID(ctx)
	userID := uuid.MustParse(payload.UserID)

	data := domain.TaskModel{
		Title:       input.Title,
		Description: &input.Description,
		Priority:    input.Priority,
		DueDate:     &dueDate,
		UserID:      userID,
	}
	errdetails := data.Validate()
	if errdetails != nil {
		return &model.TaskResponse{}, errors.New(fmt.Sprintf("%v", errdetails))
	}
	respData, err := r.Task.CreateTask(data)
	if err != nil {
		return &model.TaskResponse{}, err
	}
	taskResp := respData.Data.(resp.TaskResponse)
	return responseTaskData(taskResp), err
}

// UpdateTask is the resolver for the updateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, input model.UpdateTask) (*model.TaskResponse, error) {
	const layout = "2006-01-02 15:04:05"
	date := input.DueDate
	dueDate, err := time.Parse(layout, *date)
	if err != nil {
		return nil, fmt.Errorf("invalid duedate")
	}
	data := domain.UpdateTaskModel{
		Title:       &input.Title,
		Description: &input.Description,
		Priority:    &input.Priority,
		DueDate:     &dueDate,
		Status:      &input.Status,
	}
	taskID := uuid.MustParse(input.ID)
	errdetails := data.Validate()
	if errdetails != nil {
		return &model.TaskResponse{}, errors.New(fmt.Sprintf("%v", errdetails))
	}
	respData, err := r.Task.UpdateTask(data, taskID)
	if err != nil {
		return &model.TaskResponse{}, err
	}
	taskResp := respData.Data.(resp.TaskResponse)
	return responseTaskData(taskResp), err
}

// DeleteTask is the resolver for the deleteTask field.
func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (string, error) {
	taskID, errr := uuid.Parse(id)
	if errr != nil {
		return fmt.Sprintf("%v", errr.Error()), errr
	}
	err := r.Task.DeleteTask(taskID)
	if err.Error != nil {
		return fmt.Sprintf("Error: %v", err.Error), fmt.Errorf("%v", err.Error)
	}
	return "deleted successfully", nil
}

// GetTaskByID is the resolver for the getTaskById field.
func (r *queryResolver) GetTaskByID(ctx context.Context, id string) (*model.TaskResponse, error) {
	taskID, err := uuid.Parse(id)
	if err != nil {
		return &model.TaskResponse{}, err
	}
	data, err := r.Task.FetchTaskByID(context.Background(), taskID)
	if err != nil {
		return &model.TaskResponse{}, err
	}
	taskResp := data.Data.(resp.TaskResponse)
	return responseTaskData(taskResp), err
}

// ListTask is the resolver for the ListTask field.
func (r *queryResolver) ListTask(ctx context.Context) (*model.TaskListResponse, error) {
	data, err := r.Task.FetchTask()
	if err != nil {
		return nil, err
	}
	list := data.Data.([]resp.TaskResponse)
	return responseListTaskData(list), err
}
