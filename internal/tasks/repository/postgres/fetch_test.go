package tasks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListTask(t *testing.T) {

	for i := 0; i < 2; i++ {
		createRandomTask(t)
	}
	tasks, err := testQueries.FetchTask()
	require.NoError(t, err)
	require.NotEmpty(t, tasks)
	for _, task := range tasks {
		require.NotEmpty(t, task)

	}
}
