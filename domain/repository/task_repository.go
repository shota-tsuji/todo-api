package repository

import (
	"example.com/go-gin-todolist/domain/entity"
)

type TaskRepository interface {
	InsertTask(title string) int
	FindAllTasks() []entity.Task
}
