package main

import (
	"context"
	"example.com/go-gin-todolist/domain/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
	"os"
)

type TaskController struct {
	store *service.TaskService
}

func NewTaskController() *TaskController {
	store := service.New()
	return &TaskController{store: store}
}

func (tc *TaskController) GetTaskList(c *gin.Context) {
	allTasks := tc.store.GetAllTasks()
	c.JSON(http.StatusOK, allTasks)
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	type RequestTask struct {
		Text string `json:"text"`
	}

	var rt RequestTask
	if err := c.ShouldBindJSON(&rt); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	id := tc.store.CreateTask(rt.Text)
	c.JSON(http.StatusOK, gin.H{"Id": id})
}

func Run(lc fx.Lifecycle, controller *TaskController) {
	router := gin.Default()
	router.GET("/task/", controller.GetTaskList)
	router.POST("/task/", controller.CreateTask)
	router.Run("localhost:" + os.Getenv("SERVERPORT"))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// clean up router
			return nil
		},
	})
}

func main() {
	fx.New(
		fx.Provide(
			NewTaskController,
		),
		fx.Invoke(Run),
	).Run()
}
