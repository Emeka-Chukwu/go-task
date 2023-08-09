package auths

import (
	"database/sql"
	mockdbrep "go-task/internal/auths/repository/mock"
	"testing"

	req "go-task/domain/auths/request"
	resp "go-task/domain/auths/response"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetUserByEmail(t *testing.T) {
	userResp, _ := randomUser(t)
	testCases := []struct {
		body          req.RegisterModel
		name          string
		buildStubs    func(store *mockdbrep.MockAuthentication)
		checkResponse func(response resp.RegisterResponse, err error)
	}{
		{
			name: "Ok",
			buildStubs: func(store *mockdbrep.MockAuthentication) {

				store.EXPECT().
					GetUserByEmail(gomock.Eq(userResp.Email)).
					Times(1).
					Return(userResp.RegisterResponse, nil)
			},
			checkResponse: func(response resp.RegisterResponse, err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "InternalServerError",
			buildStubs: func(store *mockdbrep.MockAuthentication) {

				store.EXPECT().
					GetUserByEmail(gomock.Any()).
					Times(1).
					Return(resp.RegisterResponse{}, sql.ErrConnDone)
			},
			checkResponse: func(response resp.RegisterResponse, err error) {
				require.Error(t, err)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdbrep.NewMockAuthentication(ctrl)
			tc.buildStubs(store)
			authUsecase := newTestUsecase(t, store)
			loginResp, err := authUsecase.GetUserByEmail(userResp.Email)
			tc.checkResponse(loginResp, err)
		})
	}
}

func TestGetUserByID(t *testing.T) {
	userResp := randomUserId(t)
	testCases := []struct {
		body          req.RegisterModel
		name          string
		buildStubs    func(store *mockdbrep.MockAuthentication)
		checkResponse func(response resp.RegisterResponse, err error)
	}{
		{
			name: "Ok",
			buildStubs: func(store *mockdbrep.MockAuthentication) {

				store.EXPECT().
					GetUserByID(gomock.Eq(userResp.ID)).
					Times(1).
					Return(userResp, nil)
			},
			checkResponse: func(response resp.RegisterResponse, err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "InternalServerError",
			buildStubs: func(store *mockdbrep.MockAuthentication) {

				store.EXPECT().
					GetUserByID(gomock.Any()).
					Times(1).
					Return(resp.RegisterResponse{}, sql.ErrConnDone)
			},
			checkResponse: func(response resp.RegisterResponse, err error) {
				require.Error(t, err)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdbrep.NewMockAuthentication(ctrl)
			tc.buildStubs(store)
			authUsecase := newTestUsecase(t, store)
			loginResp, err := authUsecase.GetUserByID(userResp.ID)
			tc.checkResponse(loginResp, err)
		})
	}
}
