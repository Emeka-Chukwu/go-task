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

func TestUpdateTaskusercase(t *testing.T) {
	var titleEmpty = ""
	task, _ := createRandomTask(t)
	testCases := []struct {
		body          req.UpdateTaskModel
		name          string
		buildStubs    func(store *mockdb.MockTask)
		checkResponse func(response resp.TaskResponse, err error)
	}{
		{
			name: "Ok",
			body: req.UpdateTaskModel{
				Title:       &task.Title,
				Description: task.Description,
				Priority:    &task.Priority,
				DueDate:     task.DueDate,
			},
			buildStubs: func(store *mockdb.MockTask) {

				arg := req.UpdateTaskModel{
					Title:       &task.Title,
					Description: task.Description,
					Priority:    &task.Priority,
					DueDate:     task.DueDate,
					// UserID:      user.ID,
				}
				store.EXPECT().
					UpdateTask(gomock.Eq(arg), gomock.Eq(task.ID)).
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
					UpdateTask(gomock.Any(), gomock.Any()).
					Times(1).
					Return(resp.TaskResponse{}, sql.ErrConnDone)
			},
			checkResponse: func(response resp.TaskResponse, err error) {
				require.Error(t, err)
			},
			body: req.UpdateTaskModel{
				Title:       &task.Title,
				Description: task.Description,
				Priority:    &task.Priority,
				DueDate:     task.DueDate,
			},
		},

		{
			name: "RecordNotFound",
			buildStubs: func(store *mockdb.MockTask) {

				store.EXPECT().
					UpdateTask(gomock.Any(), gomock.Any()).
					Times(1).
					Return(resp.TaskResponse{}, sql.ErrNoRows)
			},
			checkResponse: func(response resp.TaskResponse, err error) {
				require.Error(t, err)
			},
			body: req.UpdateTaskModel{
				Title:       &task.Title,
				Description: task.Description,
				Priority:    &task.Priority,
				DueDate:     task.DueDate,
			},
		},
		{
			name: "InvalidData",
			buildStubs: func(store *mockdb.MockTask) {

				arg := req.TaskModel{
					Title:       "",
					Description: nil,
					Priority:    task.Priority,
					DueDate:     task.DueDate,
				}
				store.EXPECT().
					CreateTask(gomock.Eq(arg)).
					Times(0)

			},
			checkResponse: func(response resp.TaskResponse, err error) {
				require.Error(t, err)
			},
			body: req.UpdateTaskModel{
				Title:       &titleEmpty,
				Description: nil,
				Priority:    &task.Priority,
				DueDate:     task.DueDate,
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
			respData, err := authUsecase.UpdateTask(tc.body, task.ID)
			if respData.Data == nil {
				respData.Data = resp.TaskResponse{}

			}
			data := respData.Data.(resp.TaskResponse)
			tc.checkResponse(data, err)

		})
	}
}
