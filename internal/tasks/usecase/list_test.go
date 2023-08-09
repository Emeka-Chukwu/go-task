package tasks

import (
	"database/sql"
	mockdb "go-task/internal/tasks/repository/mock"

	"testing"

	req "go-task/domain/task/request"
	resp "go-task/domain/task/response"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestFetchTaskusercase(t *testing.T) {
	n := 5
	tasks := make([]resp.TaskResponse, n)
	for i := 0; i < n; i++ {
		task, _ := createRandomTask(t)
		tasks[i] = task
	}
	type taskList []resp.TaskResponse
	testCases := []struct {
		body          req.TaskModel
		name          string
		buildStubs    func(store *mockdb.MockTask)
		checkResponse func(response taskList, err error)
	}{
		{
			name: "Ok",

			buildStubs: func(store *mockdb.MockTask) {

				store.EXPECT().
					FetchTask().
					Times(1).
					Return(tasks, nil)
			},
			checkResponse: func(response taskList, err error) {
				require.NoError(t, err)
				// require.NotEmpty(t, response)
			},
		},

		{
			name: "InternalServerError",
			buildStubs: func(store *mockdb.MockTask) {

				store.EXPECT().
					FetchTask().
					Times(1).
					Return([]resp.TaskResponse{}, sql.ErrConnDone)
			},
			checkResponse: func(response taskList, err error) {
				require.Error(t, err)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockTask(ctrl)
			tc.buildStubs(store)
			authUsecase := newTestUsecase(t, store)
			respData, err := authUsecase.FetchTask()
			if respData.Data == nil {
				respData.Data = []resp.TaskResponse{}
			}
			data := respData.Data.([]resp.TaskResponse)
			tc.checkResponse(data, err)

		})
	}
}
