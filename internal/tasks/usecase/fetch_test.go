package tasks

import (
	"context"
	"database/sql"
	mockdb "go-task/internal/tasks/repository/mock"

	"testing"

	req "go-task/domain/task/request"
	resp "go-task/domain/task/response"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestFetchByTaskusercase(t *testing.T) {

	task, user := createRandomTask(t)
	testCases := []struct {
		body          req.TaskModel
		name          string
		buildStubs    func(store *mockdb.MockTask)
		checkResponse func(response resp.TaskResponse, err error)
	}{
		{
			name: "Ok",
			body: req.TaskModel{
				Title:       task.Title,
				Description: task.Description,
				Priority:    task.Priority,
				DueDate:     task.DueDate,
				// UserID:      user.ID,
			},
			buildStubs: func(store *mockdb.MockTask) {

				store.EXPECT().
					FetchTaskByID(gomock.Eq(context.Background()), gomock.Eq(task.ID)).
					Times(1).
					Return(task, nil)
			},
			checkResponse: func(response resp.TaskResponse, err error) {
				require.NoError(t, err)
				// require.NotEmpty(t, response)
			},
		},

		{
			name: "InternalServerError",
			buildStubs: func(store *mockdb.MockTask) {

				store.EXPECT().
					FetchTaskByID(gomock.Any(), gomock.Any()).
					Times(1).
					Return(resp.TaskResponse{}, sql.ErrConnDone)
			},
			checkResponse: func(response resp.TaskResponse, err error) {
				require.Error(t, err)
			},
			body: req.TaskModel{
				Title:       task.Title,
				Description: task.Description,
				Priority:    task.Priority,
				DueDate:     task.DueDate,
				UserID:      user.ID,
			},
		},
		{
			name: "RecordNotFound",
			buildStubs: func(store *mockdb.MockTask) {

				store.EXPECT().
					FetchTaskByID(gomock.Any(), gomock.Any()).
					Times(1).
					Return(resp.TaskResponse{}, sql.ErrNoRows)
			},
			checkResponse: func(response resp.TaskResponse, err error) {
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
			respData, err := authUsecase.FetchTaskByID(context.Background(), task.ID)
			if respData.Data == nil {
				respData.Data = resp.TaskResponse{}
			}
			data := respData.Data.(resp.TaskResponse)
			tc.checkResponse(data, err)

		})
	}
}
