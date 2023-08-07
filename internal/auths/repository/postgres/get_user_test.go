package auths

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetUserByID(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserByID(user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user2.Username, user1.Username)
	require.Equal(t, user2.Email, user1.Email)
	require.NotZero(t, user2.CreatedAt, user1.CreatedAt)
	require.NotZero(t, user2.UpdatedAt, user1.UpdatedAt)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second*1)
}

func TestGetUserByEmail(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserByEmail(user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user2.Username, user1.Username)
	require.Equal(t, user2.Email, user1.Email)
	require.NotZero(t, user2.CreatedAt, user1.CreatedAt)
	require.NotZero(t, user2.UpdatedAt, user1.UpdatedAt)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second*1)
}
