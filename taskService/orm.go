package taskService

type Task struct {
	ID     uint   `json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
