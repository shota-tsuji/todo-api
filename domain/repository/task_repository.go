package repository

//go:generate mockgen -source ./task_repository.go -destination task_repository_mock.generated.go -package repository

import (
	"example.com/go-gin-todolist/domain/entity"
)

type TaskRepository interface {
	InsertTask(title string) int
	FindAllTasks() []entity.Task
}
