package graph

import (
	"database/sql"
	"fmt"
	"go-task/directives"
	domain "go-task/domain/auths/request"
	respAuth "go-task/domain/auths/response"
	"go-task/middlewares"
	"go-task/token"
	"time"

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

type LoginUserGraphqlResp struct {
	CreateUser CreateUser `json:"loginUser"`
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

func TestUserLoginAccount(t *testing.T) {
	loginResp := createRandomLoginResp(t)
	user := loginResp.RegisterResponse
	testCases := []struct {
		name          string
		buildStubs    func(store *mockdb.MockAuthusecase)
		checkResponse func(response model.LoginResponse, err error)
		request       string
	}{
		{
			name: "Ok",

			buildStubs: func(store *mockdb.MockAuthusecase) {
				arg := domain.LoginModel{
					Email:    user.Email,
					Password: user.PasswordHash,
				}
				store.
					EXPECT().
					LoginUser(gomock.Eq(arg)).
					Times(1).
					Return(loginResp, nil)
			},
			checkResponse: func(response model.LoginResponse, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, response.User)
			},
			request: fmt.Sprintf(`
					mutation {
						loginUser(
						input: {email: "%s", password: "%s"}
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
			  `, user.Email, user.PasswordHash),
		},
		{
			name: "InValidData",
			buildStubs: func(store *mockdb.MockAuthusecase) {
				arg := domain.LoginModel{
					Email:    user.Email,
					Password: user.PasswordHash,
				}
				store.
					EXPECT().
					LoginUser(gomock.Eq(arg)).
					Times(0)

			},
			checkResponse: func(response model.LoginResponse, err error) {

				require.Empty(t, response.User)
			},
			request: `
			mutation {
				loginUser(
				  input: {email: "emeka233", password: "Password@"}
				) {
				  token
				  }
			  }
			  `,
		},
		{
			name: "InteranServerError",
			buildStubs: func(store *mockdb.MockAuthusecase) {
				store.
					EXPECT().
					LoginUser(gomock.Any()).
					Times(1).Return(respAuth.LoginResponse{}, sql.ErrConnDone)

			},
			checkResponse: func(response model.LoginResponse, err error) {
				require.Empty(t, response.User)
			},
			request: `
			mutation {
				loginUser(
				  input: {email: "emeka233@gmail.com", password: "Password@"}
				) {
				  token
				  }
			  }
			  `,
		},
		{
			name: "UnAuthorizedUser",
			buildStubs: func(store *mockdb.MockAuthusecase) {
				store.
					EXPECT().
					LoginUser(gomock.Any()).
					Times(1).Return(respAuth.LoginResponse{}, sql.ErrNoRows)

			},
			checkResponse: func(response model.LoginResponse, err error) {
				require.Empty(t, response.User)
			},
			request: `
			mutation {
				loginUser(
				  input: {email: "emeka233@gmail.com", password: "Password@"}
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
			var tempResp LoginUserGraphqlResp
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

func TestGetUserByIDAccount(t *testing.T) {
	loginResp := createRandomLoginResp(t)
	user := loginResp.RegisterResponse
	testCases := []struct {
		name          string
		buildStubs    func(store *mockdb.MockAuthusecase)
		checkResponse func(response model.LoginResponse, err error)
		request       string
		setupAuth     func(t *testing.T, tokenMaker token.Maker) string
	}{
		{
			name: "Ok",
			buildStubs: func(store *mockdb.MockAuthusecase) {
				store.
					EXPECT().
					GetUserByID(gomock.Eq(user.ID)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(response model.LoginResponse, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, response.User)
			},
			request: fmt.Sprintf(`
					query {
						getUserById(
						id: "%s"
						) {
							id
							username
							email
							created_at
							updated_at
						}
					}
			  `, user.ID),
			setupAuth: func(t *testing.T, tokenMaker token.Maker) string {
				return addAuthorization(t, tokenMaker, authorizationTypeBearer, user.ID.String(), time.Minute)
			},
		},
		{
			name: "InValidData",
			buildStubs: func(store *mockdb.MockAuthusecase) {
				store.
					EXPECT().
					GetUserByID(gomock.Any()).
					Times(0)

			},
			checkResponse: func(response model.LoginResponse, err error) {

				require.Empty(t, response.User)
			},
			request: `
			query {
				getUserById(
				id: "000
				) {
					id
				}
			}
			  `,
			setupAuth: func(t *testing.T, tokenMaker token.Maker) string {
				return addAuthorization(t, tokenMaker, authorizationTypeBearer, user.ID.String(), time.Minute)
			},
		},
		{
			name: "InteranServerError",
			buildStubs: func(store *mockdb.MockAuthusecase) {
				store.
					EXPECT().
					LoginUser(gomock.Any()).
					Times(1).Return(respAuth.LoginResponse{}, sql.ErrConnDone)

			},
			checkResponse: func(response model.LoginResponse, err error) {
				require.Empty(t, response.User)
			},
			request: `
			query {
				getUserById(
				 id: "000000"
				) {
				  id
				  }
			  }
			  `,
			setupAuth: func(t *testing.T, tokenMaker token.Maker) string {
				return addAuthorization(t, tokenMaker, authorizationTypeBearer, user.ID.String(), time.Minute)
			},
		},
		{
			name: "UnAuthorizedUser",
			buildStubs: func(store *mockdb.MockAuthusecase) {
				store.
					EXPECT().
					LoginUser(gomock.Any()).
					Times(1).Return(respAuth.LoginResponse{}, sql.ErrNoRows)

			},
			checkResponse: func(response model.LoginResponse, err error) {
				require.Empty(t, response.User)
				require.Error(t, err)
			},
			request: `
			query {
				loginUser(
				  input: {email: "emeka233@gmail.com", password: "Password@"}
				) {
				  token
				  }
			  }
			  `,
			setupAuth: func(t *testing.T, tokenMaker token.Maker) string {
				return ""
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			var tempResp GetUserResponse
			config, auth := newTestUsecase(t)
			config.Directives.Auth = directives.Auth
			tc.buildStubs(auth)
			srv := NewExecutableSchema(config)
			handler := handler.NewDefaultServer(srv)
			cf := util.Config{
				TokenSymmetricKey:   util.RandomString(32),
				AccessTokenDuration: time.Minute,
			}
			tokenMaker, err := token.NewJWTMaker(cf.TokenSymmetricKey)
			require.NoError(t, err)
			c := client.New(middlewares.AuthMiddleware(tokenMaker, cf, handler))
			bearer := tc.setupAuth(t, tokenMaker)
			header := client.AddHeader("Authorization", bearer)
			err = c.Post(tc.request, &tempResp, header)
			regResp := model.LoginResponse{
				User: &model.User{
					ID:       tempResp.GetUserByID.ID,
					Username: tempResp.GetUserByID.Username,
					Email:    tempResp.GetUserByID.Email,
				},
			}
			tc.checkResponse(regResp, err)
		})
	}
}

// c.Directives.Auth = directives.Auth
// 	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))
// 	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 	http.Handle("/query", middlewares.AuthMiddleware(tokenMaker, config, srv))

type GetUserResponse struct {
	GetUserByID GetUserByID `json:"getUserById"`
}

type GetUserByID struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
