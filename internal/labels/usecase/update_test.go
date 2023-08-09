package label

import (
	"database/sql"
	req "go-task/domain/label/request"
	resp "go-task/domain/label/response"
	mockdb "go-task/internal/labels/repository/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestLabelupdateusecase(t *testing.T) {
	labelresp := randomLabel(t)
	labelrespWithTask := randomLabelWithTask(t)
	// newName := util.RandomString(8)
	testCases := []struct {
		body          req.LabelModel
		name          string
		buildStubs    func(store *mockdb.MockLabel)
		checkResponse func(response resp.LabelResponse, err error)
		isWithTask    bool
	}{
		{
			name: "Ok",
			buildStubs: func(store *mockdb.MockLabel) {
				arg := req.LabelModel{
					Name: labelresp.Name,
				}
				store.EXPECT().
					Update(gomock.Eq(labelresp.ID), gomock.Eq(arg)).
					Times(1).
					Return(labelresp, nil)
			},
			checkResponse: func(response resp.LabelResponse, err error) {
				require.NoError(t, err)
			},
		},
		{
			isWithTask: true,
			name:       "OkWithtask",
			buildStubs: func(store *mockdb.MockLabel) {
				taskID := labelrespWithTask.ID.String()
				arg := req.LabelModel{
					Name:   labelresp.Name,
					TaskID: &taskID,
				}

				store.EXPECT().
					Update(gomock.Eq(labelresp.ID), gomock.Eq(arg)).
					Times(1).
					Return(labelresp, nil)

			},
			checkResponse: func(response resp.LabelResponse, err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "InternalServerError",
			buildStubs: func(store *mockdb.MockLabel) {

				store.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Times(1).
					Return(resp.LabelResponse{}, sql.ErrConnDone)
			},
			checkResponse: func(response resp.LabelResponse, err error) {
				require.Error(t, err)
			},
		},

		{
			name: "DuplicatedRecord",
			buildStubs: func(store *mockdb.MockLabel) {

				store.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Times(1).
					Return(resp.LabelResponse{}, &pq.Error{Code: "23505"})
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
			taskID := labelrespWithTask.ID.String()
			var loginResp resp.LabelResponse
			var err error
			if tc.isWithTask {
				loginResp, err = authUsecase.Update(labelresp.ID, req.LabelModel{
					Name:   labelresp.Name,
					TaskID: &taskID,
				})
			} else {
				loginResp, err = authUsecase.Update(labelresp.ID, req.LabelModel{
					Name: labelresp.Name,
				})
			}
			tc.checkResponse(loginResp, err)
		})
	}
}
