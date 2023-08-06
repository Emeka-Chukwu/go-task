package label

import (
	domain "go-task/domain/label/request"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateTaskLabel(t *testing.T) {
	createRandomTaskLabel(t)

}

func createRandomTaskLabel(t *testing.T) domain.LabelTaskModel {

	arg := domain.LabelTaskModel{
		LabelID: uuid.MustParse("e2a7945f-cf61-4d0d-959c-0850a00b5319"),
		TaskID:  uuid.MustParse("26e10ecf-c513-40f4-a0ef-b06bc9d1e2c6"),
	}
	err := testQueries.CreateTaskLabel(arg)
	require.NoError(t, err)
	return arg
}
