package userService

import "apitest/taskService"

type User struct {
	Id       uint               `json:"id"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Tasks    []taskService.Task `gorm:"foreignKey:UserID" json:"tasks,omitempty"`
}
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
