package auths

import (
	"database/sql"
	"fmt"
	mockdbrep "go-task/internal/auths/repository/mock"
	"go-task/util"
	"reflect"
	"testing"

	req "go-task/domain/auths/request"
	resp "go-task/domain/auths/response"

	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

type eqRegisterModelMatcher struct {
	arg      req.RegisterModel
	password string
}

func (e eqRegisterModelMatcher) Matches(x interface{}) bool {

	arg, ok := x.(req.RegisterModel)
	if !ok {
		return false
	}
	err := util.CheckPassword(e.password, arg.PasswordHash)
	if err != nil {
		return false
	}

	e.arg.PasswordHash = arg.PasswordHash
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqRegisterModelMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqRegisterModel(arg req.RegisterModel, password string) gomock.Matcher {
	return eqRegisterModelMatcher{arg, password}
}
func TestCreateUserusercase(t *testing.T) {
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
				arg := req.RegisterModel{
					Username: loginResp.Username,
					Email:    loginResp.Email,
				}
				store.EXPECT().
					Register(EqRegisterModel(arg, password)).
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
					Register(gomock.Any()).
					Times(1).
					Return(resp.RegisterResponse{}, sql.ErrConnDone)
			},
			checkResponse: func(response resp.LoginResponse, err error) {
				require.Error(t, err)
			},
		},
		{
			name: "DuplicatedRecord",
			buildStubs: func(store *mockdbrep.MockAuthentication) {

				store.EXPECT().
					Register(gomock.Any()).
					Times(1).
					Return(resp.RegisterResponse{}, &pq.Error{Code: "23505"})
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
			regResp, err := store.Register(req.RegisterModel{
				Username:     loginResp.Username,
				Email:        loginResp.Email,
				PasswordHash: loginResp.PasswordHash,
			})
			loginResp = resp.LoginResponse{
				RegisterResponse: regResp,
			}
			tc.checkResponse(loginResp, err)
		})
	}
}

func randomUser(t *testing.T) (user resp.LoginResponse, password string) {
	password = util.RandomString(8)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)
	user = resp.LoginResponse{
		RegisterResponse: resp.RegisterResponse{
			Username:     util.RandomUsername(),
			Email:        util.RandomEmail(),
			PasswordHash: hashedPassword,
		},
	}
	return
}
