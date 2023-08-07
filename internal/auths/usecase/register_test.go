package auths

import (
	"fmt"
	mockdbrep "go-task/internal/auths/repository/mock"
	mockdb "go-task/internal/auths/usecase/mock"
	"go-task/util"
	"reflect"
	"testing"

	req "go-task/domain/auths/request"
	resp "go-task/domain/auths/response"

	"github.com/golang/mock/gomock"
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
		buildStubs    func(storeusecase *mockdb.MockAuthusecase, storeRep *mockdbrep.MockAuthentication)
		checkResponse func(response resp.LoginResponse, err error)
	}{
		{
			name: "Ok",
			buildStubs: func(storeuse *mockdb.MockAuthusecase, storerep *mockdbrep.MockAuthentication) {
				arg := req.RegisterModel{
					Username: loginResp.Username,
					Email:    loginResp.Email,
				}
				storeuse.EXPECT().
					Register(EqRegisterModel(arg, password)).
					Times(1).
					Return(loginResp, nil)
			},
			checkResponse: func(response resp.LoginResponse, err error) {
				require.NoError(t, err)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			storeRep := mockdbrep.NewMockAuthentication(ctrl)
			store := mockdb.NewMockAuthusecase(ctrl)
			tc.buildStubs(store, storeRep)
			loginResp, err := store.Register(req.RegisterModel{
				Username:     loginResp.Username,
				Email:        loginResp.Email,
				PasswordHash: loginResp.PasswordHash,
			})
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
