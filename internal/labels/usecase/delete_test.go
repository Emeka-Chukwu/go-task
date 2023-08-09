package label

import (
	"database/sql"
	mockdb "go-task/internal/labels/repository/mock"

	"testing"

	req "go-task/domain/label/request"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestDeleteLabelrusercase(t *testing.T) {
	labelresp := randomLabel(t)
	testCases := []struct {
		body          req.LabelModel
		name          string
		buildStubs    func(store *mockdb.MockLabel)
		checkResponse func(err error)
	}{
		{
			name: "Ok",
			buildStubs: func(store *mockdb.MockLabel) {

				store.EXPECT().
					Delete(gomock.Eq(labelresp.ID)).
					Times(1).
					Return(nil)
			},
			checkResponse: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "InternalServerError",
			buildStubs: func(store *mockdb.MockLabel) {
				store.EXPECT().
					Delete(gomock.Any()).
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
			store := mockdb.NewMockLabel(ctrl)
			tc.buildStubs(store)
			authUsecase := newTestUsecase(t, store)
			err := authUsecase.Delete(labelresp.ID)
			tc.checkResponse(err)
		})
	}
}
