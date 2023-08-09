package label

import (
	"database/sql"
	req "go-task/domain/label/request"
	resp "go-task/domain/label/response"
	mockdb "go-task/internal/labels/repository/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetLabelTasksByID(t *testing.T) {
	labelResp := randomLabel(t)

	testCases := []struct {
		body          req.LabelModel
		name          string
		buildStubs    func(store *mockdb.MockLabel)
		checkResponse func(response resp.LabelResponse, err error)
	}{
		{
			name: "Ok",
			buildStubs: func(store *mockdb.MockLabel) {
				store.EXPECT().
					GetByID(gomock.Eq(labelResp.ID)).
					Times(1).
					Return(labelResp, nil)
			},
			checkResponse: func(response resp.LabelResponse, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, response)
			},
		},

		{
			name: "InternalServerError",
			buildStubs: func(store *mockdb.MockLabel) {

				store.EXPECT().
					GetByID(gomock.Any()).
					Times(1).
					Return(resp.LabelResponse{}, sql.ErrConnDone)
			},
			checkResponse: func(response resp.LabelResponse, err error) {
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
			loginResp, err := authUsecase.GetByID(labelResp.ID)
			tc.checkResponse(loginResp, err)
		})
	}
}
