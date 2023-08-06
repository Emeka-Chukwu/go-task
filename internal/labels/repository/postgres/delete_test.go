package label

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeleteLabel(t *testing.T) {
	label1 := createRandomLabel(t)
	err := testQueries.Delete(label1.ID)
	require.NoError(t, err)
	label2, err := testQueries.GetByID(label1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, label2)
}
