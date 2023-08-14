package graph

import (
	"context"
	domain "go-task/domain/auths/request"
	respAuth "go-task/domain/auths/response"
	"go-task/graph/model"
	mockdb "go-task/internal/auths/usecase/mock"
	"go-task/util"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUserCreationAccount(t *testing.T) {
	loginResp := createRandomLoginResp(t)
	user := loginResp.RegisterResponse
	testCases := []struct {
		body          model.NewUser
		name          string
		buildStubs    func(store *mockdb.MockAuthusecase)
		checkResponse func(response model.LoginResponse, err error)
	}{
		{
			name: "Ok",
			body: model.NewUser{
				Username:     user.Username,
				Email:        user.Email,
				PasswordHash: user.PasswordHash,
			},
			buildStubs: func(store *mockdb.MockAuthusecase) {
				arg := domain.RegisterModel{
					Username:     user.Username,
					Email:        user.Email,
					PasswordHash: user.PasswordHash,
				}
				store.
					EXPECT().
					Register(gomock.Eq(arg)).
					Times(2).
					Return(loginResp, nil)
			},
			checkResponse: func(response model.LoginResponse, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, response.User)
			},
		},
		{
			name: "InValidData",
			body: model.NewUser{
				Username:     user.Username,
				Email:        "emmem",
				PasswordHash: user.PasswordHash,
			},
			buildStubs: func(store *mockdb.MockAuthusecase) {
				arg := domain.RegisterModel{
					Username:     user.Username,
					Email:        user.Email,
					PasswordHash: user.PasswordHash,
				}
				store.
					EXPECT().
					Register(gomock.Eq(arg)).
					Times(0)

			},
			checkResponse: func(response model.LoginResponse, err error) {
				require.Error(t, err)
				require.Empty(t, response.User)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockAuthusecase(ctrl)
			tc.buildStubs(store)
			authgql := newTestUsecase(t)
			loginResp, err := authgql.CreateUser(context.Background(), tc.body)
			tc.checkResponse(*loginResp, err)
		})
	}
}

func createRandomLoginResp(t *testing.T) respAuth.LoginResponse {
	user := respAuth.RegisterResponse{
		ID:           util.Getuuid(),
		Username:     util.RandomUsername(),
		PasswordHash: util.RandomPassword(),
		Email:        util.RandomEmail(),
	}
	loginResp := respAuth.LoginResponse{
		RegisterResponse: user,
	}
	return loginResp
}
