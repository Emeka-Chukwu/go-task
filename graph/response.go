package graph

import (
	resp "go-task/domain/auths/response"

	respLabel "go-task/domain/label/response"
	respTask "go-task/domain/task/response"
	"go-task/graph/model"
)

func responseUser(user resp.LoginResponse) *model.LoginResponse {
	return &model.LoginResponse{
		User: &model.User{
			ID:        user.ID.String(),
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
		Token:     user.Token,
		ExpiredAt: user.ExpiredAt.String(),
	}
}

func responseUserData(user resp.RegisterResponse) *model.User {
	return &model.User{
		ID:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}

func responseTaskData(task respTask.TaskResponse) *model.TaskResponse {
	date := task.DueDate.String()
	return &model.TaskResponse{
		Message: "Task fetched successfully",
		Data: &model.Task{
			ID:          task.ID.String(),
			Title:       task.Title,
			Description: *task.Description,
			DueDate:     &date,
			Priority:    task.Priority,
			CreatedAt:   task.CreatedAt.GoString(),
			UpdatedAt:   task.UpdatedAt.GoString(),
			Status:      task.Status,
		},
	}
}

func responseListTaskData(tasks []respTask.TaskResponse) *model.TaskListResponse {
	tasksResp := make([]*model.Task, 0)
	for _, task := range tasks {
		date := task.DueDate.String()
		taskData := &model.Task{
			ID:          task.ID.String(),
			Title:       task.Title,
			Description: *task.Description,
			DueDate:     &date,
			Priority:    task.Priority,
			CreatedAt:   task.CreatedAt.GoString(),
			UpdatedAt:   task.UpdatedAt.GoString(),
			Status:      task.Status,
		}
		tasksResp = append(tasksResp, taskData)
	}
	return &model.TaskListResponse{
		Message: "Task fetched successfully",
		Data:    tasksResp,
	}
}

// = responseListTaskData(tempResp.Lists.Data)

func responseLabelData(label respLabel.LabelResponse) *model.LabelResponse {
	return &model.LabelResponse{
		Message: "Label fetched successfully",
		Data: &model.Label{
			ID:        label.ID.String(),
			Name:      label.Name,
			CreatedAt: label.CreatedAt.GoString(),
			UpdatedAt: label.UpdatedAt.GoString(),
		},
	}
}

func responseLabelListData(labels []respLabel.LabelResponse) *model.LabelListResponse {
	respLabels := make([]*model.Label, 0)
	for _, label := range labels {
		labelresp := model.Label{
			ID:        label.ID.String(),
			Name:      label.Name,
			CreatedAt: label.CreatedAt.GoString(),
			UpdatedAt: label.UpdatedAt.GoString(),
		}

		respLabels = append(respLabels, &labelresp)
	}
	return &model.LabelListResponse{
		Data:    respLabels,
		Message: "Label fetch successfully",
	}
}

func responseLabelTask(labels []respLabel.LabelTaskResponse) []*model.LabelTaskResponse {
	respList := make([]*model.LabelTaskResponse, 0)
	for _, resp := range labels {

		ltResp := model.LabelTaskResponse{
			Label: &model.Label{
				ID:        resp.ID.String(),
				Name:      resp.Name,
				CreatedAt: resp.CreatedAt.GoString(),
				UpdatedAt: resp.UpdatedAt.GoString(),
			},
			////// always change this Task to any
			Task: resp.Tasks,
		}
		respList = append(respList, &ltResp)
	}
	return respList
}

// type LabelTaskResponse struct {
//     Label *Label  `json:"label"`
//     Task  []*Task `json:"task,omitempty"`
// }
