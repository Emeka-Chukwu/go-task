package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"go-task/graph/model"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.TaskResponse, error) {
	panic(fmt.Errorf("not implemented: CreateTask - createTask"))
}

// UpdateTask is the resolver for the updateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, input model.UpdateTask) (*model.TaskResponse, error) {
	panic(fmt.Errorf("not implemented: UpdateTask - updateTask"))
}

// DeleteTask is the resolver for the deleteTask field.
func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented: DeleteTask - deleteTask"))
}

// GetTaskByID is the resolver for the getTaskById field.
func (r *queryResolver) GetTaskByID(ctx context.Context, id string) (*model.Task, error) {
	panic(fmt.Errorf("not implemented: GetTaskByID - getTaskById"))
}

// ListTask is the resolver for the ListTask field.
func (r *queryResolver) ListTask(ctx context.Context) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented: ListTask - ListTask"))
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) Delete(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented: Delete - delete"))
}