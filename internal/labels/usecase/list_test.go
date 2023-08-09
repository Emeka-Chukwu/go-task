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

func TestGetLabelListusercase(t *testing.T) {
	n := 5
	labels := make([]resp.LabelResponse, n)
	for i := 0; i < n; i++ {
		labels[i] = randomLabel(t)
	}
	type LabelResponse []resp.LabelResponse
	testCases := []struct {
		body          req.LabelModel
		name          string
		buildStubs    func(store *mockdb.MockLabel)
		checkResponse func(response LabelResponse, err error)
	}{
		{
			name: "Ok",
			buildStubs: func(store *mockdb.MockLabel) {
				store.EXPECT().
					List().
					Times(1).
					Return(labels, nil)
			},
			checkResponse: func(response LabelResponse, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, response)
			},
		},

		{
			name: "InternalServerError",
			buildStubs: func(store *mockdb.MockLabel) {

				store.EXPECT().
					List().
					Times(1).
					Return([]resp.LabelResponse{}, sql.ErrConnDone)
			},
			checkResponse: func(response LabelResponse, err error) {
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
			loginResp, err := authUsecase.List()
			tc.checkResponse(loginResp, err)
		})
	}
}
