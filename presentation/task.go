package presentation

import (
	"example.com/go-gin-todolist/domain/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskController struct {
	taskService *service.TaskService
}

func NewTaskController(taskService *service.TaskService) *TaskController {
	return &TaskController{taskService: taskService}
}

// GetTaskList godoc
// @Summary      List tasks
// @Description  get tasks
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Success      200  {array} []entity.Task
// @Router /tasks [get]
func (tc *TaskController) GetTaskList(c *gin.Context) {
	allTasks := tc.taskService.GetAllTasks()
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

	id := tc.taskService.CreateTask(rt.Text)
	c.JSON(http.StatusOK, gin.H{"Id": id})
}
