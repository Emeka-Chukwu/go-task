package graph

import (
	"database/sql"
	"fmt"
	domain "go-task/domain/auths/request"
	respAuth "go-task/domain/auths/response"

	// "go-task/graph/generated"

	"go-task/graph/model"
	mockdb "go-task/internal/auths/usecase/mock"
	"go-task/util"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/golang/mock/gomock"

	// "github.com/mrdulin/gqlgen-cnode/graph/model"
	"github.com/stretchr/testify/require"
)

func TestUserCreationAccount(t *testing.T) {
	loginResp := createRandomLoginResp(t)
	user := loginResp.RegisterResponse
	testCases := []struct {
		body          model.NewUser
		name          string
		buildStubs    func(store *mockdb.MockAuthusecase)
		checkResponse func(response model.LoginResponse, err error)
		request       string
	}{
		{
			name: "Ok",
			body: model.NewUser{
				Username:     user.Username,
				Email:        user.Email,
				PasswordHash: user.PasswordHash,
			},
			buildStubs: func(store *mockdb.MockAuthusecase) {
				arg := domain.RegisterModel{
					Username:     user.Username,
					Email:        user.Email,
					PasswordHash: user.PasswordHash,
				}
				store.
					EXPECT().
					Register(gomock.Eq(arg)).
					Times(1).
					Return(loginResp, nil)
			},
			checkResponse: func(response model.LoginResponse, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, response.User)
			},
			request: fmt.Sprintf(`
					mutation {
						createUser(
						input: {username: "%s", email: "%s", passwordHash: "%s"}
						) {
						token
						expired_at
						user {
							id
							username
							email
							created_at
							updated_at
						}
						}
					}
			  `, user.Username, user.Email, user.PasswordHash),
		},
		{
			name: "InValidData",
			body: model.NewUser{
				Username:     user.Username,
				Email:        "emmem",
				PasswordHash: user.PasswordHash,
			},
			buildStubs: func(store *mockdb.MockAuthusecase) {
				arg := domain.RegisterModel{
					Username:     user.Username,
					Email:        user.Email,
					PasswordHash: user.PasswordHash,
				}
				store.
					EXPECT().
					Register(gomock.Eq(arg)).
					Times(0)

			},
			checkResponse: func(response model.LoginResponse, err error) {

				require.Empty(t, response.User)
			},
			request: ` 
			mutation {
				createUser(
				  input: {username: "emeka pas", email: "emeka233", passwordHash: "Password@"}
				) {
				  token
				  }
			  }
			  `,
		},
		{
			name: "InteranServerError",
			body: model.NewUser{
				Username:     user.Username,
				Email:        user.Email,
				PasswordHash: user.PasswordHash,
			},
			buildStubs: func(store *mockdb.MockAuthusecase) {
				store.
					EXPECT().
					Register(gomock.Any()).
					Times(1).Return(respAuth.LoginResponse{}, sql.ErrConnDone)

			},
			checkResponse: func(response model.LoginResponse, err error) {

				require.Empty(t, response.User)
			},
			request: ` 
			mutation {
				createUser(
				  input: {username: "emeka pas", email: "emeka233@gmail.com", passwordHash: "Password@"}
				) {
				  token
				  }
			  }
			  `,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			var tempResp LoginGraphqlResp
			config, auth := newTestUsecase(t)
			tc.buildStubs(auth)
			srv := NewExecutableSchema(config)
			c := client.New(handler.NewDefaultServer(srv))
			err := c.Post(tc.request, &tempResp)
			regResp := model.LoginResponse{
				Token: tempResp.CreateUser.Token,
				User: &model.User{
					ID:       tempResp.CreateUser.User.ID,
					Username: tempResp.CreateUser.User.Username,
					Email:    tempResp.CreateUser.User.Email,
				},
			}
			tc.checkResponse(regResp, err)

		})
	}
}

func createRandomLoginResp(t *testing.T) respAuth.LoginResponse {
	user := respAuth.RegisterResponse{
		ID:           util.Getuuid(),
		Username:     util.RandomUsername(),
		PasswordHash: util.RandomPassword(),
		Email:        util.RandomEmail(),
	}
	loginResp := respAuth.LoginResponse{
		RegisterResponse: user,
	}
	return loginResp
}

type LoginGraphqlResp struct {
	CreateUser CreateUser `json:"createUser"`
}

type CreateUser struct {
	Token     string `json:"token"`
	ExpiredAt string `json:"expired_at"`
	User      User   `json:"user"`
}

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

// func TestUserLoginAccount(t *testing.T) {
// 	loginResp := createRandomLoginResp(t)
// 	user := loginResp.RegisterResponse
// 	testCases := []struct {
// 		name          string
// 		buildStubs    func(store *mockdb.MockAuthusecase)
// 		checkResponse func(response model.LoginResponse, err error)
// 		request       string
// 	}{
// 		{
// 			name: "Ok",

// 			buildStubs: func(store *mockdb.MockAuthusecase) {
// 				arg := domain.LoginModel{
// 					Email:    user.Email,
// 					Password: user.PasswordHash,
// 				}
// 				store.
// 					EXPECT().
// 					Register(gomock.Eq(arg)).
// 					Times(1).
// 					Return(loginResp, nil)
// 			},
// 			checkResponse: func(response model.LoginResponse, err error) {
// 				require.NoError(t, err)
// 				require.NotEmpty(t, response.User)
// 			},
// 			request: fmt.Sprintf(`
// 					mutation {
// 						createUser(
// 						input: {username: "%s", email: "%s", passwordHash: "%s"}
// 						) {
// 						token
// 						expired_at
// 						user {
// 							id
// 							username
// 							email
// 							created_at
// 							updated_at
// 						}
// 						}
// 					}
// 			  `, user.Username, user.Email, user.PasswordHash),
// 		},
// 		{
// 			name: "InValidData",
// 			buildStubs: func(store *mockdb.MockAuthusecase) {
// 				arg := domain.RegisterModel{
// 					Username:     user.Username,
// 					Email:        user.Email,
// 					PasswordHash: user.PasswordHash,
// 				}
// 				store.
// 					EXPECT().
// 					Register(gomock.Eq(arg)).
// 					Times(0)

// 			},
// 			checkResponse: func(response model.LoginResponse, err error) {

// 				require.Empty(t, response.User)
// 			},
// 			request: `
// 			mutation {
// 				createUser(
// 				  input: {username: "emeka pas", email: "emeka233", passwordHash: "Password@"}
// 				) {
// 				  token
// 				  }
// 			  }
// 			  `,
// 		},
// 		{
// 			name: "InteranServerError",
// 			buildStubs: func(store *mockdb.MockAuthusecase) {
// 				store.
// 					EXPECT().
// 					Register(gomock.Any()).
// 					Times(1).Return(respAuth.LoginResponse{}, sql.ErrConnDone)

// 			},
// 			checkResponse: func(response model.LoginResponse, err error) {

// 				require.Empty(t, response.User)
// 			},
// 			request: `
// 			mutation {
// 				createUser(
// 				  input: {username: "emeka pas", email: "emeka233@gmail.com", passwordHash: "Password@"}
// 				) {
// 				  token
// 				  }
// 			  }
// 			  `,
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]
// 		t.Run(tc.name, func(t *testing.T) {
// 			var tempResp LoginGraphqlResp
// 			config, auth := newTestUsecase(t)
// 			tc.buildStubs(auth)
// 			srv := NewExecutableSchema(config)
// 			c := client.New(handler.NewDefaultServer(srv))
// 			err := c.Post(tc.request, &tempResp)
// 			regResp := model.LoginResponse{
// 				Token: tempResp.CreateUser.Token,
// 				User: &model.User{
// 					ID:       tempResp.CreateUser.User.ID,
// 					Username: tempResp.CreateUser.User.Username,
// 					Email:    tempResp.CreateUser.User.Email,
// 				},
// 			}
// 			tc.checkResponse(regResp, err)

// 		})
// 	}
// }
