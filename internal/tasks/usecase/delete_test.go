package tasks

import (
	"database/sql"
	mockdb "go-task/internal/tasks/repository/mock"

	"testing"

	req "go-task/domain/task/request"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestDeleteTaskusercase(t *testing.T) {
	task, _ := createRandomTask(t)
	testCases := []struct {
		body          req.TaskModel
		name          string
		buildStubs    func(store *mockdb.MockTask)
		checkResponse func(err error)
	}{
		{
			name: "Ok",
			body: req.TaskModel{
				Title:       task.Title,
				Description: task.Description,
				Priority:    task.Priority,
				DueDate:     task.DueDate,
			},
			buildStubs: func(store *mockdb.MockTask) {
				store.EXPECT().
					DeleteTask(gomock.Eq(task.ID)).
					Times(1).
					Return(nil)
			},
			checkResponse: func(err error) {
				require.NoError(t, err)
				// require.NotEmpty(t, response)
			},
		},

		{
			name: "InternalServerError",
			buildStubs: func(store *mockdb.MockTask) {

				store.EXPECT().
					DeleteTask(gomock.Any()).
					Times(1).
					Return(sql.ErrConnDone)
			},
			checkResponse: func(err error) {
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

			respData := authUsecase.DeleteTask(task.ID)
			if respData.Error == nil {
				tc.checkResponse(nil)
			} else {
				data := respData.Error.(error)
				tc.checkResponse(data)
			}

		})
	}
}
