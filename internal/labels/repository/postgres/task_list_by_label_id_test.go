package label

import (
	domain "go-task/domain/label/request"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTaskListByLabelID(t *testing.T) {
	var lasttl domain.LabelTaskModel
	for i := 0; i < 2; i++ {
		lasttl = createRandomTaskLabel(t)
	}
	ltResp, err := testQueries.ListByLabelID(lasttl.LabelID)
	require.NoError(t, err)
	require.NotZero(t, ltResp)
	require.NotEmpty(t, ltResp)

}
