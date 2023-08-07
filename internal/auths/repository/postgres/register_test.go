package auths

import (
	domain "go-task/domain/auths/request"
	resp "go-task/domain/auths/response"
	"go-task/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) resp.RegisterResponse {
	arg := domain.RegisterModel{
		Username:     util.RandomUsername(),
		Email:        util.RandomEmail(),
		PasswordHash: util.RandomPassword(),
	}
	user, err := testQueries.Register(arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.Email, arg.Email)
	require.NotZero(t, user.CreatedAt, user.UpdatedAt)
	return user
}

func TestRegisterUser(t *testing.T) {
	createRandomUser(t)
}
