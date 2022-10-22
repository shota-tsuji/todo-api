package main

import (
	"example.com/go-gin-todolist/domain/repository"
	"example.com/go-gin-todolist/domain/service"
	"example.com/go-gin-todolist/infrastructure/mysql"
	"example.com/go-gin-todolist/presentation"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
	"net/http"
	"os"
)

func Run(controller *presentation.TaskController) {
	router := gin.Default()
	router.Handle(http.MethodGet, "/task/", controller.GetTaskList)
	router.Handle(http.MethodPost, "/task/", controller.CreateTask)
	err := router.Run("localhost:" + os.Getenv("SERVERPORT"))
	if err != nil {
		log.Fatal("server failed.")
		return
	}
}

func main() {
	fx.New(
		fx.Provide(
			fx.Annotate(mysql.NewRepository, fx.As(new(repository.TaskRepository))),
			mysql.NewMysqlSession,
			service.NewTaskService,
			presentation.NewTaskController,
		),
		fx.Invoke(Run),
	).Run()
}
