package label

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTasksByLabel(t *testing.T) {
	for i := 0; i < 2; i++ {
		createRandomTaskLabel(t)
	}
	ltResp, err := testQueries.ListByLabel()
	require.NoError(t, err)
	for _, resp := range ltResp {
		require.NotEmpty(t, resp)
	}
}
