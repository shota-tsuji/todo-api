package main

import (
	"example.com/go-gin-todolist/internal/taskstore"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type taskServer struct {
	store *taskstore.TaskStore
}

func NewTaskServer() *taskServer {
	store := taskstore.New()
	return &taskServer{store: store}
}

func (ts *taskServer) getTaskList(c *gin.Context) {
	allTasks := ts.store.GetAllTasks()
	c.JSON(http.StatusOK, allTasks)
}

func (ts *taskServer) createTask(c *gin.Context) {
	type RequestTask struct {
		Text string `json:"text"`
	}

	var rt RequestTask
	if err := c.ShouldBindJSON(&rt); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	id := ts.store.CreateTask(rt.Text)
	c.JSON(http.StatusOK, gin.H{"Id": id})
}

func main() {
	router := gin.Default()
	server := NewTaskServer()

	router.GET("/task/", server.getTaskList)
	router.POST("/task/", server.createTask)

	router.Run("localhost:" + os.Getenv("SERVERPORT"))
}
