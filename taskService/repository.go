package taskService

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTasksByUserID(userID uint) ([]Task, error)
	UpdateTaskByID(id uint, updatedtask Task) (Task, error)
	DeleteTaskByID(id int) error
}
type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}
func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}
func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var task []Task
	err := r.db.Find(&task).Error
	return task, err
}
func (r *taskRepository) GetTasksByUserID(userID uint) ([]Task, error) {
	var tasks []Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, updatedtask Task) (Task, error) {
	var task Task
	first := r.db.First(&task, id)
	if first.Error != nil {
		if errors.Is(first.Error, gorm.ErrRecordNotFound) {
			return Task{}, errors.New("Task not found")
		}
		log.Printf("Error finding task: %v", first.Error)
		return Task{}, fmt.Errorf("Error finding task: %w", first.Error)
	}
	result := r.db.Model(&task).Where("id = ?", id).Updates(updatedtask)
	if result.Error != nil {
		log.Println("bad repository")
		return Task{}, result.Error
	}
	return task, nil
}
func (r *taskRepository) DeleteTaskByID(id int) error {
	result := r.db.Where("id = ?", id).Delete(&Task{})
	if result.Error != nil {
		log.Println("bad repository")
		return result.Error
	}
	return nil
}
