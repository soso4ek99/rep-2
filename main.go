package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

var task string

type TaskPayload struct {
	Task string `json:"task"`
}

func Posthandler(c echo.Context) error {
	payload := new(TaskPayload)
	if err := c.Bind(payload); err != nil {
		fmt.Println("ошибка добавление task", err)
		return c.String(http.StatusBadRequest, "Invalid JSON payload")
	}
	task = payload.Task
	fmt.Printf("Task '%s' получено и сохранено.", task)
	return c.String(http.StatusOK, fmt.Sprintf("Task получен: %s", task))
}
func gethandler(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("hello, %s", task))
}

func main() {
	e := echo.New()

	e.POST("/post", Posthandler)
	e.GET("/get", gethandler)
	fmt.Println("server запущен на: 8080")
	e.Start(":8080")
}
