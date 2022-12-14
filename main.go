package main

import (
	"example.com/todo-api/config"
	_ "example.com/todo-api/docs"
	"example.com/todo-api/domain/repository"
	"example.com/todo-api/domain/service"
	"example.com/todo-api/infrastructure/mysql"
	"example.com/todo-api/presentation"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"log"
	"net/http"
)

// @title Task Management API
// @version 1.0
// @description This is a task management application.

// @host localhost:8080
// @BasePath  /api/v1

func Run(controller *presentation.TaskController, sc config.ServerConfig) {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.Handle(http.MethodGet, "", controller.GetTaskList)
			tasks.Handle(http.MethodPost, "", controller.CreateTask)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(sc.Host + ":" + sc.Port)
	if err != nil {
		log.Fatal("server failed.")
		return
	}
}

func main() {
	fx.New(
		config.Module,
		fx.Provide(
			fx.Annotate(mysql.NewRepository, fx.As(new(repository.TaskRepository))),
			mysql.NewMysqlSession,
			service.NewTaskService,
			presentation.NewTaskController,
		),
		fx.Invoke(Run),
	).Run()
}
