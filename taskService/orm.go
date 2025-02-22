package taskService

type Task struct {
	ID     uint   `json:"id"`
	Text   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

var tasks []Task

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var nextID = 1
