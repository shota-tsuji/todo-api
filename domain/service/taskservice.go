package service

import (
	"example.com/go-gin-todolist/domain/entity"
	"example.com/go-gin-todolist/domain/repository"
)

type TaskService struct {
	repository repository.TaskRepository
}

func NewTaskService(repository repository.TaskRepository) *TaskService {
	return &TaskService{repository: repository}
}

func (ts *TaskService) GetAllTasks() []entity.Task {
	return ts.repository.FindAllTasks()
}

func (ts *TaskService) CreateTask(title string) int {
	return ts.repository.InsertTask(title)
}
