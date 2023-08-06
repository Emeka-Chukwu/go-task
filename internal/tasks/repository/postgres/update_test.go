package tasks

import (
	domain "go-task/domain/task/request"
	"go-task/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestUpdateTask(t *testing.T) {
	oldTask := createRandomTask(t)
	newTitle := util.RandomString(8)
	newDescription := util.RandomString(8)
	newPriority := util.RandomString(8)
	newDueDate := time.Now().Add(time.Hour * 24)
	status := util.RandomString(8)
	arg := domain.UpdateTaskModel{
		Title:       &newTitle,
		Description: &newDescription,
		Priority:    &newPriority,
		DueDate:     &newDueDate,
		Status:      &status,
	}
	updatedTask, err := testQueries.UpdateTask(arg, oldTask.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedTask)
	require.NotEqual(t, oldTask.Title, updatedTask.Title)
	require.Equal(t, oldTask.ID, updatedTask.ID)
	require.NotZero(t, updatedTask.ID)
	require.WithinDuration(t, oldTask.CreatedAt, updatedTask.CreatedAt, time.Second*2)
}
