package main

import (
	"apitest/internal/database"
	"apitest/internal/handlers"
	"apitest/internal/web/tasks"
	"apitest/internal/web/users"
	"apitest/taskService"
	"apitest/userService"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()

	userrepo := userService.NewUserRepository(database.DB)
	userservice := userService.NewUserService(userrepo)
	userhandler := handlers.NewHandlers(userservice)
	taskrepo := taskService.NewTaskRepository(database.DB)
	taskservice := taskService.NewService(taskrepo)
	taskhandler := handlers.NewHandler(taskservice)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	userstrictHandler := users.NewStrictHandler(userhandler, nil)
	users.RegisterHandlers(e, userstrictHandler)
	strictHandler := tasks.NewStrictHandler(taskhandler, nil)
	tasks.RegisterHandlers(e, strictHandler)
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
	fmt.Println("test new")
}
