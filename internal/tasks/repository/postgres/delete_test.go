package tasks

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeleteTask(t *testing.T) {
	task1 := createRandomTask(t)
	err := testQueries.DeleteTask(task1.ID)
	require.NoError(t, err)
	task2, err := testQueries.FetchTaskByID(context.Background(), task1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, task2)
}
