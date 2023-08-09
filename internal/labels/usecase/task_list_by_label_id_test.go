package label

import (
	"database/sql"
	mockdb "go-task/internal/labels/repository/mock"

	"testing"

	req "go-task/domain/label/request"
	resp "go-task/domain/label/response"

	respTask "go-task/domain/task/response"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetLabelListBuLabelusecase(t *testing.T) {

	ltResp := createRandomLabelTask(t)
	testCases := []struct {
		body          req.LabelModel
		name          string
		buildStubs    func(store *mockdb.MockLabel)
		checkResponse func(response resp.LabelTaskResponse, err error)
	}{
		{
			name: "Ok",
			buildStubs: func(store *mockdb.MockLabel) {
				store.EXPECT().
					ListByLabelID(gomock.Eq(ltResp.ID)).
					Times(1).
					Return(ltResp, nil)
			},
			checkResponse: func(response resp.LabelTaskResponse, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, response)
			},
		},

		{
			name: "InternalServerError",
			buildStubs: func(store *mockdb.MockLabel) {

				store.EXPECT().
					ListByLabelID(gomock.Eq(ltResp.ID)).
					Times(1).
					Return(resp.LabelTaskResponse{}, sql.ErrConnDone)
			},
			checkResponse: func(response resp.LabelTaskResponse, err error) {
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
			loginResp, err := authUsecase.ListByLabelID(ltResp.ID)
			tc.checkResponse(loginResp, err)
		})
	}
}

func createRandomTaskLabel(t *testing.T) (task respTask.TaskResponse, label resp.LabelResponse) {
	task = randomLabelWithTask(t)
	label = randomLabel(t)
	return

}

func createRandomLabelTask(t *testing.T) (ltResp resp.LabelTaskResponse) {
	ltResp = resp.LabelTaskResponse{
		LabelResponse: randomLabel(t),
	}
	return
}
