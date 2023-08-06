package label

import (
	domain "go-task/domain/label/request"
	"go-task/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestUpdateLabel(t *testing.T) {
	oldLabel := createRandomLabel(t)
	oldName := util.RandomUsername()
	arg := domain.LabelModel{
		Name: oldName,
	}
	newLabel, err := testQueries.Update(oldLabel.ID, arg)
	require.NoError(t, err)
	require.NotEmpty(t, newLabel)
	require.NotEqual(t, oldLabel.Name, newLabel.Name)
	require.NotZero(t, oldLabel.ID, newLabel.ID)
	require.WithinDuration(t, oldLabel.UpdatedAt, newLabel.UpdatedAt, time.Second*3)
}
