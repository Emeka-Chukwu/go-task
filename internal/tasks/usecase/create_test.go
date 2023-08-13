package tasks

import (
	"database/sql"
	"fmt"
	mockdb "go-task/internal/tasks/repository/mock"
	"time"

	"go-task/util"
	"testing"

	domain "go-task/domain/auths/response"
	req "go-task/domain/task/request"
	resp "go-task/domain/task/response"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateTaskusercase(t *testing.T) {

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

				arg := req.TaskModel{
					Title:       task.Title,
					Description: task.Description,
					Priority:    task.Priority,
					DueDate:     task.DueDate,
					// UserID:      user.ID,
				}
				store.EXPECT().
					CreateTask(gomock.Eq(arg)).
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

				// arg := req.TaskModel{
				// 	Title:       util.RandomString(8),
				// 	Description: task.Description,
				// 	Priority:    task.Priority,
				// 	DueDate:     task.DueDate,
				// 	UserID:      user.ID,
				// }
				store.EXPECT().
					CreateTask(gomock.Any()).
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
			name: "InvalidData",
			buildStubs: func(store *mockdb.MockTask) {

				arg := req.TaskModel{
					Title:       "",
					Description: nil,
					Priority:    task.Priority,
					DueDate:     task.DueDate,
					UserID:      user.ID,
				}
				store.EXPECT().
					CreateTask(gomock.Eq(arg)).
					Times(0)

			},
			checkResponse: func(response resp.TaskResponse, err error) {
				require.Error(t, err)
			},
			body: req.TaskModel{
				Title:       "",
				Description: nil,
				Priority:    task.Priority,
				DueDate:     task.DueDate,
				UserID:      user.ID,
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
			respData, err := authUsecase.CreateTask(tc.body)
			if respData.Data == nil {
				respData.Data = resp.TaskResponse{}
			}
			data := respData.Data.(resp.TaskResponse)
			tc.checkResponse(data, err)

		})
	}
}

func createRandomTask(t *testing.T) (task resp.TaskResponse, user domain.RegisterResponse) {
	var priority = "high"
	var dueDate = time.Now().AddDate(0, 0, 2).Local()
	fmt.Println(dueDate)
	description := util.RandomString(30)

	user = domain.RegisterResponse{
		ID:           util.Getuuid(),
		Username:     util.RandomUsername(),
		Email:        util.RandomEmail(),
		PasswordHash: util.RandomPassword(),
	}
	task = resp.TaskResponse{
		// ID:          util.Getuuid(),
		Title:       util.RandomString(8),
		Description: &description,
		Priority:    priority,
		DueDate:     &dueDate,
	}

	user = domain.RegisterResponse{
		ID:           util.Getuuid(),
		Username:     util.RandomUsername(),
		Email:        util.RandomEmail(),
		PasswordHash: util.RandomPassword(),
	}
	return
}
