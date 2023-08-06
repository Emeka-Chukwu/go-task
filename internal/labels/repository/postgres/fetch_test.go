package label

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetLabelById(t *testing.T) {
	label1 := createRandomLabel(t)
	label2, err := testQueries.GetByID(label1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, label2)
	require.Equal(t, label1.Name, label2.Name)
	require.NotZero(t, label2.CreatedAt)
	require.Equal(t, label2.CreatedAt, label1.CreatedAt)
	require.NotZero(t, label2.CreatedAt)
	require.WithinDuration(t, label1.CreatedAt, label2.CreatedAt, time.Second*2)
}
