package service

import (
	"example.com/todo-api/domain/entity"
	"example.com/todo-api/domain/repository"
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
