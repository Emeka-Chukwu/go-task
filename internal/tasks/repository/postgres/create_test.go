package tasks

import (
	"testing"
	"time"

	domain "go-task/domain/task/request"
	resp "go-task/domain/task/response"
	"go-task/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomTask(t *testing.T) resp.TaskResponse {
	var priority = "high"
	var dueDate = time.Now().Add(time.Hour * 24)
	description := util.RandomString(30)
	uuid := uuid.MustParse("cd06dc60-1daa-41c5-a22c-e2f8141a7bcd")
	arg := domain.TaskModel{
		Title:       util.RandomString(8),
		Description: &description,
		Priority:    priority,
		DueDate:     &dueDate,
		UserID:      uuid,
	}
	task, err := testQueries.CreateTask(arg)
	require.NoError(t, err)
	require.NotEmpty(t, task)
	require.Equal(t, arg.Title, task.Title)
	require.Equal(t, arg.Title, task.Title)
	require.Equal(t, arg.Description, task.Description)
	require.Equal(t, arg.Priority, task.Priority)
	require.NotZero(t, task.CreatedAt)
	require.NotZero(t, task.UpdatedAt)
	require.Equal(t, task.Status, "todo")
	return task
}

func TestCreateTask(t *testing.T) {
	createRandomTask(t)
}
