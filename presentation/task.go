package presentation

import (
	"example.com/go-gin-todolist/domain/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
