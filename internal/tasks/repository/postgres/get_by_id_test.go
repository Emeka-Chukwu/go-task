package tasks

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetTaskByID(t *testing.T) {
	task1 := createRandomTask(t)
	task2, err := testQueries.FetchTaskByID(context.Background(), task1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, task2)
	require.Equal(t, task1.Title, task2.Title)
	require.Equal(t, task1.ID, task2.ID)
	require.NotZero(t, task1.ID)
	require.WithinDuration(t, task1.CreatedAt, task2.CreatedAt, time.Second*2)

}
