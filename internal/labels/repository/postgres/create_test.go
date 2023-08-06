package label

import (
	"testing"

	domain "go-task/domain/label/request"
	resp "go-task/domain/label/response"
	"go-task/util"

	"github.com/stretchr/testify/require"
)

func createRandomLabel(t *testing.T) resp.LabelResponse {

	arg := domain.LabelModel{
		Name: util.RandomString(8),
	}
	label, err := testQueries.Create(arg)
	require.NoError(t, err)
	require.NotEmpty(t, label)
	require.Equal(t, arg.Name, label.Name)
	require.NotZero(t, label.CreatedAt)
	require.NotZero(t, label.UpdatedAt)

	return label
}

func TestCreateLabel(t *testing.T) {
	createRandomLabel(t)
}
