package label

import (
	"database/sql"
	mockdb "go-task/internal/labels/repository/mock"

	"testing"

	req "go-task/domain/label/request"
	resp "go-task/domain/label/response"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetLabelListByLabelusecase(t *testing.T) {
	n := 5
	ltResps := make([]resp.LabelTaskResponse, n)
	for i := 0; i < n; i++ {
		ltResps[i] = createRandomLabelTask(t)
	}
	type listTaskResponse []resp.LabelTaskResponse
	testCases := []struct {
		body          req.LabelModel
		name          string
		buildStubs    func(store *mockdb.MockLabel)
		checkResponse func(response listTaskResponse, err error)
	}{
		{
			name: "Ok",
			buildStubs: func(store *mockdb.MockLabel) {
				store.EXPECT().
					ListByLabel().
					Times(1).
					Return(ltResps, nil)
			},
			checkResponse: func(response listTaskResponse, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, response)
			},
		},
		{
			name: "InternalServerError",
			buildStubs: func(store *mockdb.MockLabel) {

				store.EXPECT().
					ListByLabel().
					Times(1).
					Return(ltResps, sql.ErrConnDone)
			},
			checkResponse: func(response listTaskResponse, err error) {
				require.Error(t, err)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockLabel(ctrl)
			tc.buildStubs(store)
			authUsecase := newTestUsecase(t, store)
			loginResp, err := authUsecase.ListByLabel()
			tc.checkResponse(loginResp, err)
		})
	}
}
