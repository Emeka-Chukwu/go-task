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

func TestLoginUserusercase(t *testing.T) {
	loginResp, password := randomUser(t)
	testCases := []struct {
		body          req.RegisterModel
		name          string
		buildStubs    func(store *mockdbrep.MockAuthentication)
		checkResponse func(response resp.LoginResponse, err error)
	}{
		{
			name: "Ok",
			buildStubs: func(store *mockdbrep.MockAuthentication) {
				arg := req.LoginModel{
					Password: password,
					Email:    loginResp.Email,
				}
				store.EXPECT().
					GetUserByEmail(gomock.Eq(arg.Email)).
					Times(1).
					Return(loginResp.RegisterResponse, nil)
			},
			checkResponse: func(response resp.LoginResponse, err error) {
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
			checkResponse: func(response resp.LoginResponse, err error) {
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
			loginResp, err := authUsecase.LoginUser(req.LoginModel{Email: loginResp.Email, Password: password})
			tc.checkResponse(loginResp, err)
		})
	}
}
