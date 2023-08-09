package label

import (
	"database/sql"
	mockdb "go-task/internal/labels/repository/mock"
	"time"

	"go-task/util"
	"testing"

	req "go-task/domain/label/request"
	resp "go-task/domain/label/response"
	respTask "go-task/domain/task/response"

	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestCreateLabelusercase(t *testing.T) {
	labelresp := randomLabel(t)
	labelrespWithTask := randomLabelWithTask(t)
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
					Create(gomock.Eq(arg)).
					Times(1).
					Return(labelresp, nil)
			},
			checkResponse: func(response resp.LabelResponse, err error) {
				require.NoError(t, err)
			},
			isWithTask: false,
		},
		{
			name: "OkWithtask",
			buildStubs: func(store *mockdb.MockLabel) {
				taskID := labelrespWithTask.ID.String()
				arg := req.LabelModel{
					Name:   labelresp.Name,
					TaskID: &taskID,
				}

				store.EXPECT().
					Create(gomock.Eq(arg)).
					Times(1).
					Return(labelresp, nil)

			},
			checkResponse: func(response resp.LabelResponse, err error) {
				require.NoError(t, err)
			},
			isWithTask: true,
		},
		{
			name: "InternalServerError",
			buildStubs: func(store *mockdb.MockLabel) {

				store.EXPECT().
					Create(gomock.Any()).
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
					Create(gomock.Any()).
					Times(1).
					Return(resp.LabelResponse{}, &pq.Error{Code: "23505"})
			},
			checkResponse: func(response resp.LabelResponse, err error) {
				require.Error(t, err)
			},
		},
		{
			name: "DuplicatedRecord",
			buildStubs: func(store *mockdb.MockLabel) {

				store.EXPECT().
					Create(gomock.Any()).
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
				loginResp, err = authUsecase.Create(req.LabelModel{
					Name:   labelresp.Name,
					TaskID: &taskID,
				})
			} else {
				loginResp, err = authUsecase.Create(req.LabelModel{
					Name: labelresp.Name,
				})
			}
			tc.checkResponse(loginResp, err)
		})
	}
}

func randomLabel(t *testing.T) (user resp.LabelResponse) {
	user = resp.LabelResponse{
		ID:   util.Getuuid(),
		Name: util.RandomString(8),
	}
	return
}
func randomLabelWithTask(t *testing.T) (task respTask.TaskResponse) {
	var priority = "high"
	var dueDate = time.Now().Add(time.Hour * 24)
	description := util.RandomString(30)
	task = respTask.TaskResponse{
		ID:          util.Getuuid(),
		Title:       util.RandomString(8),
		Description: &description,
		Priority:    priority,
		DueDate:     &dueDate,
	}
	return
}
