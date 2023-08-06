package label

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListLabel(t *testing.T) {
	// doma
	for i := 0; i < 2; i++ {
		createRandomLabel(t)
	}
	labels, err := testQueries.List()
	require.NoError(t, err)
	for _, label := range labels {
		require.NotEmpty(t, label)
	}
}
