package tasks

// FetchTask implements Taskusecase.
func (t *taskusecase) FetchTask() (ResponseData, error) {
	resp, err := t.store.FetchTask()
	if err != nil {
		return ResponseData{Message: "Internal server err"}, err
	}
	return ResponseData{Message: "Data fetched successfully", Data: resp}, nil
}
