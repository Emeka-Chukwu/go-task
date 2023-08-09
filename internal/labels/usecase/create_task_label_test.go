package label

import (
	"database/sql"
	mockdb "go-task/internal/labels/repository/mock"

	"testing"

	req "go-task/domain/label/request"

	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestCreateTaskLabelusercase(t *testing.T) {
	labelresp := randomLabel(t)
	task := randomLabelWithTask(t)
	testCases := []struct {
		body          req.LabelTaskModel
		name          string
		buildStubs    func(store *mockdb.MockLabel)
		checkResponse func(err error)
	}{
		{
			name: "Ok",
			buildStubs: func(store *mockdb.MockLabel) {
				arg := req.LabelTaskModel{
					TaskID:  task.ID,
					LabelID: labelresp.ID,
				}
				store.EXPECT().
					CreateTaskLabel(gomock.Eq(arg)).
					Times(1).
					Return(nil)
			},
			checkResponse: func(err error) {
				require.NoError(t, err)
			},
			body: req.LabelTaskModel{
				TaskID:  task.ID,
				LabelID: labelresp.ID,
			},
		},

		{
			name: "InternalServerError",
			buildStubs: func(store *mockdb.MockLabel) {

				store.EXPECT().
					CreateTaskLabel(gomock.Any()).
					Times(1).
					Return(sql.ErrConnDone)
			},
			checkResponse: func(err error) {
				require.Error(t, err)
			},
		},
		{
			name: "DuplicatedRecord",
			buildStubs: func(store *mockdb.MockLabel) {

				store.EXPECT().
					CreateTaskLabel(gomock.Any()).
					Times(1).
					Return(&pq.Error{Code: "23505"})
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
			store := mockdb.NewMockLabel(ctrl)
			tc.buildStubs(store)
			authUsecase := newTestUsecase(t, store)
			err := authUsecase.CreateTaskLabel(tc.body)

			tc.checkResponse(err)
		})
	}
}

// func randomLabel(t *testing.T) (user resp.LabelResponse) {
// 	user = resp.LabelResponse{
// 		ID:   util.Getuuid(),
// 		Name: util.RandomString(8),
// 	}
// 	return
// }
// func randomLabelWithTask(t *testing.T) (task respTask.TaskResponse) {
// 	var priority = "high"
// 	var dueDate = time.Now().Add(time.Hour * 24)
// 	description := util.RandomString(30)
// 	task = respTask.TaskResponse{
// 		ID:          util.Getuuid(),
// 		Title:       util.RandomString(8),
// 		Description: &description,
// 		Priority:    priority,
// 		DueDate:     &dueDate,
// 	}
// 	return
// }
