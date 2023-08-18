package graph

import (
	"context"
	"database/sql"
	"fmt"
	"go-task/directives"
	domain "go-task/domain/task/request"
	respTask "go-task/domain/task/response"
	"go-task/middlewares"
	"go-task/token"
	"strings"
	"time"

	// "go-task/graph/generated"

	"go-task/graph/model"
	tasks "go-task/internal/tasks/usecase"
	mockdb "go-task/internal/tasks/usecase/mock"
	"go-task/util"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/golang/mock/gomock"

	// "github.com/mrdulin/gqlgen-cnode/graph/model"
	"github.com/stretchr/testify/require"
)

func TestCreateTaskAccount(t *testing.T) {
	taskResp := createRandomTaskResop(t)
	user := createRandomLoginResp(t)
	respData := tasks.ResponseData{
		Message: "Success",
		Data:    taskResp,
	}

	date1 := strings.Split(taskResp.DueDate.String(), " +")
	date := date1[0]
	testCases := []struct {
		name          string
		buildStubs    func(store *mockdb.MockTaskusecase)
		checkResponse func(response model.TaskResponse, err error)
		request       string

		setupAuth func(t *testing.T, tokenMaker token.Maker) string
	}{
		{
			name: "Ok",
			buildStubs: func(store *mockdb.MockTaskusecase) {
				arg := domain.TaskModel{
					Title:       taskResp.Title,
					Description: taskResp.Description,
					Priority:    taskResp.Priority,
					DueDate:     taskResp.DueDate,
					UserID:      user.ID,
				}
				store.
					EXPECT().
					CreateTask(gomock.Eq(arg)).
					Times(1).
					Return(respData, nil)
			},
			checkResponse: func(response model.TaskResponse, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, response.Data)
			},
			request: fmt.Sprintf(`
					mutation {
						createTask(
						input: {title: "%s", description: "%s", priority: "%s", due_date: "%s"}
						) {
							Data{
								id
								title
								description
								due_date
								status
								priority
								user_id
								created_at
								updated_at
							  }
						}
					}
			  `, taskResp.Title, *taskResp.Description, taskResp.Priority, date),
			setupAuth: func(t *testing.T, tokenMaker token.Maker) string {
				return addAuthorization(t, tokenMaker, authorizationTypeBearer, user.ID.String(), time.Minute)
			},
		},
		{
			name: "InValidData",

			buildStubs: func(store *mockdb.MockTaskusecase) {

				store.
					EXPECT().
					CreateTask(gomock.Any()).
					Times(0)

			},
			checkResponse: func(response model.TaskResponse, err error) {

			},
			request: fmt.Sprintf(`
			mutation {
				createTask(
				input: {title: "%s", description: "%s", priority: "%s", due_date: "%s"}
				) {
					Data{
						id
					  }
				}
			}
			  `, taskResp.Title, *taskResp.Description, taskResp.Priority, taskResp.DueDate),
			setupAuth: func(t *testing.T, tokenMaker token.Maker) string {
				return addAuthorization(t, tokenMaker, authorizationTypeBearer, user.ID.String(), time.Minute)
			},
		},
		{
			name: "InValidDataSecond",

			buildStubs: func(store *mockdb.MockTaskusecase) {

				store.
					EXPECT().
					CreateTask(gomock.Any()).
					Times(0)

			},
			checkResponse: func(response model.TaskResponse, err error) {

			},
			request: fmt.Sprintf(`
			mutation {
				createTask(
				input: {title: "%s", description: "%s", priority: "%s", due_date: "%s"}
				) {
					Data{
						id
					  }
				}
			}
			  `, taskResp.Title, *taskResp.Description, "uuuu", date),
			setupAuth: func(t *testing.T, tokenMaker token.Maker) string {
				return addAuthorization(t, tokenMaker, authorizationTypeBearer, user.ID.String(), time.Minute)
			},
		},
		{
			name: "InteranServerError",
			setupAuth: func(t *testing.T, tokenMaker token.Maker) string {
				return addAuthorization(t, tokenMaker, authorizationTypeBearer, user.ID.String(), time.Minute)
			},
			buildStubs: func(store *mockdb.MockTaskusecase) {
				store.
					EXPECT().
					CreateTask(gomock.Any()).
					Times(1).Return(tasks.ResponseData{}, sql.ErrConnDone)
			},
			checkResponse: func(response model.TaskResponse, err error) {
				require.Error(t, err)
				require.Error(t, err)
				require.Empty(t, response.Data.Title)
				require.Empty(t, response.Data.ID)
				require.Empty(t, response.Data.Status)
			},
			request: fmt.Sprintf(`
					mutation {
						createTask(
						input: {title: "%s", description: "%s", priority: "%s", due_date: "%s"}
						) {
							Data{
								id
								title
								description
								due_date
								status
								priority
								user_id
								created_at
								updated_at
							  }
						}
					}
			  `, taskResp.Title, *taskResp.Description, taskResp.Priority, date),
		},
		{
			name: "UnAuthorizedUser",
			buildStubs: func(store *mockdb.MockTaskusecase) {
				store.
					EXPECT().
					CreateTask(gomock.Any()).
					Times(0)

			},
			checkResponse: func(response model.TaskResponse, err error) {
				require.Empty(t, response.Data.Title)
				require.Empty(t, response.Data.ID)
				require.Error(t, err)
			},
			request: `
			query {
				createTask(
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
			var tempResp CreateTaskRespons
			config, task := newTestUsecase(t)
			config.Directives.Auth = directives.Auth
			tc.buildStubs(task.Task)
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
			taskResp := model.TaskResponse{
				Data: &model.Task{
					ID:          tempResp.CreateTask.Data.ID,
					Title:       tempResp.CreateTask.Data.Title,
					Description: tempResp.CreateTask.Data.Description,
					DueDate:     &tempResp.CreateTask.Data.DueDate,
					Status:      tempResp.CreateTask.Data.Status,
					CreatedAt:   tempResp.CreateTask.Data.CreatedAt,
					UpdatedAt:   tempResp.CreateTask.Data.UpdatedAt,
				},
			}
			tc.checkResponse(taskResp, err)
		})
	}
}

func TestUpdateTaskAccount(t *testing.T) {
	taskResp := createRandomTaskResop(t)
	user := createRandomLoginResp(t)
	respData := tasks.ResponseData{
		Message: "Success",
		Data:    taskResp,
	}

	date1 := strings.Split(taskResp.DueDate.String(), " +")
	date := date1[0]
	testCases := []struct {
		name          string
		buildStubs    func(store *mockdb.MockTaskusecase)
		checkResponse func(response model.TaskResponse, err error)
		request       string
		setupAuth     func(t *testing.T, tokenMaker token.Maker) string
	}{
		{
			name: "Ok",
			buildStubs: func(store *mockdb.MockTaskusecase) {
				arg := domain.UpdateTaskModel{
					Title:       &taskResp.Title,
					Description: taskResp.Description,
					Priority:    &taskResp.Priority,
					DueDate:     taskResp.DueDate,
					Status:      &taskResp.Status,
				}
				store.
					EXPECT().
					UpdateTask(gomock.Eq(arg), gomock.Eq(taskResp.ID)).
					Times(1).
					Return(respData, nil)
			},
			checkResponse: func(response model.TaskResponse, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, response.Data)
			},
			request: fmt.Sprintf(`
					mutation {
						updateTask(
						input: {title: "%s", description: "%s", priority: "%s", due_date: "%s", status: "%s", id:"%s"}
						) {
							Data{
								id
								title
								description
								due_date
								status
								priority
								user_id
								created_at
								updated_at
							  }
						}
					}
			  `, taskResp.Title, *taskResp.Description, taskResp.Priority, date, "in-progress", taskResp.ID),
			setupAuth: func(t *testing.T, tokenMaker token.Maker) string {
				return addAuthorization(t, tokenMaker, authorizationTypeBearer, user.ID.String(), time.Minute)
			},
		},
		{
			name: "InValidData",

			buildStubs: func(store *mockdb.MockTaskusecase) {

				store.
					EXPECT().
					UpdateTask(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(response model.TaskResponse, err error) {
			},
			request: fmt.Sprintf(`
			mutation {
				updateTask(
				input: {title: "%s", description: "%s", priority: "%s", due_date: "%s"}
				) {
					Data{
						id
					  }
				}
			}
			  `, taskResp.Title, *taskResp.Description, taskResp.Priority, taskResp.DueDate),
			setupAuth: func(t *testing.T, tokenMaker token.Maker) string {
				return addAuthorization(t, tokenMaker, authorizationTypeBearer, user.ID.String(), time.Minute)
			},
		},
		{
			name: "InValidDataSecond",

			buildStubs: func(store *mockdb.MockTaskusecase) {

				store.
					EXPECT().
					UpdateTask(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(response model.TaskResponse, err error) {
				require.Error(t, err)
			},
			request: fmt.Sprintf(`
			mutation {
				updateTask(
				input: {title: "%s", description: "%s", priority: "%s", due_date: "%s",id: "%s", status:"in-progresso"}
				) {
					Data{
						id
					  }
				}
			}
			  `, taskResp.Title, *taskResp.Description, "uuuu", date, taskResp.ID),
			setupAuth: func(t *testing.T, tokenMaker token.Maker) string {
				return addAuthorization(t, tokenMaker, authorizationTypeBearer, user.ID.String(), time.Minute)
			},
		},
		{
			name: "InteranServerError",
			setupAuth: func(t *testing.T, tokenMaker token.Maker) string {
				return addAuthorization(t, tokenMaker, authorizationTypeBearer, user.ID.String(), time.Minute)
			},
			buildStubs: func(store *mockdb.MockTaskusecase) {
				store.
					EXPECT().
					UpdateTask(gomock.Any(), gomock.Any()).
					Times(1).Return(tasks.ResponseData{}, sql.ErrConnDone)
			},
			checkResponse: func(response model.TaskResponse, err error) {
				require.Error(t, err)
				require.Error(t, err)
				require.Empty(t, response.Data.Title)
				require.Empty(t, response.Data.ID)
				require.Empty(t, response.Data.Status)
			},
			request: fmt.Sprintf(`
					mutation {
						updateTask(
						input: {title: "%s", description: "%s", priority: "%s", due_date: "%s", status:"in-progress", id:"%s"}
						) {
							Data{
								id
								title
								description
								due_date
								status
								priority
								user_id
								created_at
								updated_at
							  }
						}
					}
			  `, taskResp.Title, *taskResp.Description, taskResp.Priority, date, taskResp.ID),
		},
		{
			name: "UnAuthorizedUser",
			buildStubs: func(store *mockdb.MockTaskusecase) {
				store.
					EXPECT().
					UpdateTask(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(response model.TaskResponse, err error) {
				require.Empty(t, response.Data.Title)
				require.Empty(t, response.Data.ID)
				require.Error(t, err)
			},
			request: `
			query {
				updateUser(
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
			var tempResp UpdateTaskRespons
			config, task := newTestUsecase(t)
			config.Directives.Auth = directives.Auth
			tc.buildStubs(task.Task)
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
			taskResp := model.TaskResponse{
				Data: &model.Task{
					ID:          tempResp.CreateTask.Data.ID,
					Title:       tempResp.CreateTask.Data.Title,
					Description: tempResp.CreateTask.Data.Description,
					DueDate:     &tempResp.CreateTask.Data.DueDate,
					Status:      tempResp.CreateTask.Data.Status,
					CreatedAt:   tempResp.CreateTask.Data.CreatedAt,
					UpdatedAt:   tempResp.CreateTask.Data.UpdatedAt,
				},
			}
			tc.checkResponse(taskResp, err)
		})
	}
}

func TestGetTaskByID(t *testing.T) {
	taskResp := createRandomTaskResop(t)
	user := createRandomLoginResp(t)
	respData := tasks.ResponseData{
		Message: "Success",
		Data:    taskResp,
	}
	testCases := []struct {
		name          string
		buildStubs    func(store *mockdb.MockTaskusecase)
		checkResponse func(response model.TaskResponse, err error)
		request       string
		setupAuth     func(t *testing.T, tokenMaker token.Maker) string
	}{
		{
			name: "Ok",
			buildStubs: func(store *mockdb.MockTaskusecase) {

				store.
					EXPECT().
					FetchTaskByID(gomock.Eq(context.Background()), gomock.Eq(taskResp.ID)).
					Times(1).
					Return(respData, nil)
			},
			checkResponse: func(response model.TaskResponse, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, response.Data)
			},
			request: fmt.Sprintf(`
					query {
						getTaskById (
						id: "%s"
						) {
							Data{
								id
								title
								description
								due_date
								status
								priority
								user_id
								created_at
								updated_at
							  }
						}
					}
			  `, taskResp.ID),
			setupAuth: func(t *testing.T, tokenMaker token.Maker) string {
				return addAuthorization(t, tokenMaker, authorizationTypeBearer, user.ID.String(), time.Minute)
			},
		},
		{
			name: "InValidData",

			buildStubs: func(store *mockdb.MockTaskusecase) {

				store.
					EXPECT().
					FetchTaskByID(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(response model.TaskResponse, err error) {
				require.Error(t, err)
			},
			request: fmt.Sprintf(`
			query {
				getTaskById(
				id: "%s"
				) {
					Data{
						id
					  }
				}
			}
			  `, "983090389"),
			setupAuth: func(t *testing.T, tokenMaker token.Maker) string {
				return addAuthorization(t, tokenMaker, authorizationTypeBearer, user.ID.String(), time.Minute)
			},
		},
		{
			name: "RecordNotFound",

			buildStubs: func(store *mockdb.MockTaskusecase) {

				store.
					EXPECT().
					FetchTaskByID(gomock.Any(), gomock.Any()).
					Times(1).
					Return(tasks.ResponseData{}, sql.ErrNoRows)

			},
			checkResponse: func(response model.TaskResponse, err error) {
				require.Error(t, err)
			},
			request: fmt.Sprintf(`
			query {
				getTaskById(
				id: "%s"
				) {
					Data{
						id
					  }
				}
			}
			  `, util.Getuuid().String()),
			setupAuth: func(t *testing.T, tokenMaker token.Maker) string {
				return addAuthorization(t, tokenMaker, authorizationTypeBearer, user.ID.String(), time.Minute)
			},
		},
		{
			name: "InteranServerError",
			setupAuth: func(t *testing.T, tokenMaker token.Maker) string {
				return addAuthorization(t, tokenMaker, authorizationTypeBearer, user.ID.String(), time.Minute)
			},
			buildStubs: func(store *mockdb.MockTaskusecase) {
				store.
					EXPECT().
					FetchTaskByID(gomock.Any(), gomock.Any()).
					Times(1).Return(tasks.ResponseData{}, sql.ErrConnDone)
			},
			checkResponse: func(response model.TaskResponse, err error) {
				require.Error(t, err)
				require.Error(t, err)
				require.Empty(t, response.Data.Title)
				require.Empty(t, response.Data.ID)
				require.Empty(t, response.Data.Status)
			},
			request: fmt.Sprintf(`
					query {
						getTaskById(
					id: "%s"
						) {
							Data{
								id
								title
								description
								due_date
								status
								priority
								user_id
								created_at
								updated_at
							  }
						}
					}
			  `, taskResp.ID),
		},
		{
			name: "UnAuthorizedUser",
			buildStubs: func(store *mockdb.MockTaskusecase) {
				store.
					EXPECT().
					FetchTaskByID(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(response model.TaskResponse, err error) {
				require.Empty(t, response.Data.Title)
				require.Empty(t, response.Data.ID)
				require.Error(t, err)
			},
			request: `
			query {
				getTaskById (
				  id:""
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
			var tempResp GetTaskByIdRespons
			config, task := newTestUsecase(t)
			config.Directives.Auth = directives.Auth
			tc.buildStubs(task.Task)
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
			taskResp := model.TaskResponse{
				Data: &model.Task{
					ID:          tempResp.CreateTask.Data.ID,
					Title:       tempResp.CreateTask.Data.Title,
					Description: tempResp.CreateTask.Data.Description,
					DueDate:     &tempResp.CreateTask.Data.DueDate,
					Status:      tempResp.CreateTask.Data.Status,
					CreatedAt:   tempResp.CreateTask.Data.CreatedAt,
					UpdatedAt:   tempResp.CreateTask.Data.UpdatedAt,
				},
			}
			tc.checkResponse(taskResp, err)
		})
	}
}

func TestFetchTasks(t *testing.T) {
	n := 3
	var tasksList = []respTask.TaskResponse{}
	for i := 0; i < n; i++ {
		taskResp := createRandomTaskResop(t)
		tasksList = append(tasksList, taskResp)
	}
	user := createRandomLoginResp(t)
	respData := tasks.ResponseData{
		Message: "Success",
		Data:    tasksList,
	}
	testCases := []struct {
		name          string
		buildStubs    func(store *mockdb.MockTaskusecase)
		checkResponse func(response model.TaskListResponse, err error)
		request       string
		setupAuth     func(t *testing.T, tokenMaker token.Maker) string
	}{
		{
			name: "Ok",
			buildStubs: func(store *mockdb.MockTaskusecase) {
				store.
					EXPECT().
					FetchTask().
					Times(1).
					Return(respData, nil)
			},
			checkResponse: func(response model.TaskListResponse, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, response.Data)
			},
			request: `
					query {
						ListTask {
							Data{
								id
								title
								description
								due_date
								status
								priority
								user_id
								created_at
								updated_at
							  }
						}
					}
			  `,
			setupAuth: func(t *testing.T, tokenMaker token.Maker) string {
				return addAuthorization(t, tokenMaker, authorizationTypeBearer, user.ID.String(), time.Minute)
			},
		},

		{
			name: "InteranServerError",
			setupAuth: func(t *testing.T, tokenMaker token.Maker) string {
				return addAuthorization(t, tokenMaker, authorizationTypeBearer, user.ID.String(), time.Minute)
			},
			buildStubs: func(store *mockdb.MockTaskusecase) {
				store.
					EXPECT().
					FetchTask().
					Times(1).Return(tasks.ResponseData{}, sql.ErrConnDone)
			},
			checkResponse: func(response model.TaskListResponse, err error) {
				require.Error(t, err)
				require.Error(t, err)
				for _, value := range response.Data {
					require.Empty(t, value.CreatedAt)
					require.Empty(t, value.ID)
					require.Empty(t, value.Priority)
				}

			},
			request: `
					query {
						ListTask {
							Data{
								id
								title
								description
								due_date
								status
								priority
								user_id
								created_at
								updated_at
							  }
						}
					}
			  `,
		},
		{
			name: "UnAuthorizedUser",
			buildStubs: func(store *mockdb.MockTaskusecase) {
				store.
					EXPECT().
					FetchTaskByID(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(response model.TaskListResponse, err error) {
				for _, value := range response.Data {
					require.Empty(t, value.CreatedAt)
					require.Empty(t, value.ID)
					require.Empty(t, value.Priority)
				}
				require.Error(t, err)
			},
			request: `
			query {
				ListTask  {
				  id
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
			var tempResp ListTasksResponse
			config, task := newTestUsecase(t)
			config.Directives.Auth = directives.Auth
			tc.buildStubs(task.Task)
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
			tasksResp := make([]*model.Task, 0)
			for _, task := range tempResp.Lists.Data {
				date := task.DueDate
				taskData := &model.Task{
					ID:          task.ID,
					Title:       task.Title,
					Description: task.Description,
					DueDate:     &date,
					Priority:    task.Priority,
					CreatedAt:   task.CreatedAt,
					UpdatedAt:   task.UpdatedAt,
					Status:      task.Status,
				}
				tasksResp = append(tasksResp, taskData)
			}
			data := &model.TaskListResponse{
				Message: "Task fetched successfully",
				Data:    tasksResp,
			}
			tc.checkResponse(*data, err)
		})
	}
}

// ////////////
func createRandomTaskResop(t *testing.T) respTask.TaskResponse {
	description := util.RandomString(5)
	dueDate := time.Now().Add(time.Hour * 24).UTC()
	task := respTask.TaskResponse{
		ID:          util.Getuuid(),
		Title:       util.RandomString(16),
		Description: &description,
		Status:      "in-progress",
		Priority:    "high",
		DueDate:     &dueDate,
	}

	return task
}

type CreateTaskRespons struct {
	CreateTask CreateTask `json:"createTask"`
}
type UpdateTaskRespons struct {
	CreateTask CreateTask `json:"updateTask"`
}
type GetTaskByIdRespons struct {
	CreateTask CreateTask `json:"getTaskById"`
}

type ListTasksResponse struct {
	Lists ListTask `json:"listTask"`
}

// updateTask
type CreateTask struct {
	Data Data `json:"Data"`
}
type UpdateTask struct {
	Data Data `json:"Data"`
}
type ListTask struct {
	Data []Data `json:"Data"`
}

type Data struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	UserID      string `json:"user_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
