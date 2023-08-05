package tasks

import resp "go-task/domain/task/response"

// Fetch implements Task.
func (*task) FetchTask() ([]*resp.TaskResponse, error) {
	panic("unimplemented")
}
