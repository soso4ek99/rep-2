package taskService

type Task struct {
	ID     uint   `json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserID uint   `json:"user_id"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
