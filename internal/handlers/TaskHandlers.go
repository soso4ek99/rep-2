package handlers

import (
	"apitest/internal/web/tasks"
	"apitest/taskService"
	"context"
	"log"
)

type Handler struct {
	Service *taskService.TaskService
}

func (h *Handler) GetUsersIdTasks(ctx context.Context, request tasks.GetUsersIdTasksRequestObject) (tasks.GetUsersIdTasksResponseObject, error) {
	userID := request.Id
	tasksForUser, err := h.Service.GetTasksByUserID(userID)
	if err != nil {
		return nil, err
	}

	if len(tasksForUser) == 0 {
		return tasks.GetUsersIdTasks200JSONResponse{}, nil
	}

	var response tasks.GetUsersIdTasks200JSONResponse
	for _, tsk := range tasksForUser {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := int(request.Id) // Преобразуем int в uint

	// 2. Удаление задачи из сервиса
	err := h.Service.DeleteTask(id)
	if err != nil {

		return tasks.DeleteTasksId500JSONResponse{}, err // 500
	}
	// 3. Возврат успешного ответа (204 No Content)
	return tasks.DeleteTasksId204Response{}, nil
}

func (h *Handler) PutTasksId(ctx context.Context, request tasks.PutTasksIdRequestObject) (tasks.PutTasksIdResponseObject, error) {
	id := uint(request.Id)
	body := request.Body
	if body == nil {
		log.Println("Error: Request body is nil")
		return tasks.PutTasksId400JSONResponse{}, nil // 400 Bad Request
	}

	// Проверяем на nil и используем значения по умолчанию, если они не указаны
	taskValue := ""
	if body.Task != nil {
		taskValue = *body.Task
	}
	isDoneValue := false // Значение по умолчанию для isDone
	if body.IsDone != nil {
		isDoneValue = *body.IsDone
	}

	// Создаем структуру Task для передачи в сервис
	updatedTask := taskService.Task{
		ID:     id,          // Важно: передаем ID в структуру Task
		Task:   taskValue,   // Описание задачи
		IsDone: isDoneValue, // Статус выполнения
	}

	// *Предполагаем*, что UpdateTask принимает ID и структуру Task и возвращает Task и error
	upTask, err := h.Service.UpdateTask(id, updatedTask) // передаем всю структуру, а не параметры по отдельности
	if err != nil {
		return tasks.PutTasksId500JSONResponse{}, err // 500 Internal Server Error
	}

	response := tasks.PutTasksId200JSONResponse{
		Id:     &upTask.ID,     // Возвращаем ID обновленной задачи
		Task:   &upTask.Task,   // Возвращаем описание обновленной задачи
		IsDone: &upTask.IsDone, // Возвращаем статус обновленной задачи
	}
	return response, nil // 200 OK
}
func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		log.Println("Error getting all tasks:", err)
		return tasks.GetTasks500JSONResponse{}, err // 500
	}

	// Преобразуем задачи из сервиса в формат, ожидаемый API
	response := tasks.GetTasks200JSONResponse{}
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID, // Преобразуем uint в *uint64
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}

	return response, nil // 200 OK
}
func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// 1. Извлечение тела запроса
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: taskRequest.IsDone,
		UserID: *taskRequest.UserId,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate) // Создаем задачу через сервис
	if err != nil {
		log.Println("Error creating task:", err)
		return tasks.PostTasks500JSONResponse{}, err // 500
	}

	// 4. Формирование успешного ответа (201 Created)
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserID,
	}
	return response, nil // 201 Created
}
func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}
