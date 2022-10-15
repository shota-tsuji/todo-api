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

func (ts *taskServer) getAllTasksHandler(c *gin.Context) {
	allTasks := ts.store.GetAllTasks()
	c.JSON(http.StatusOK, allTasks)
}

func main() {
	router := gin.Default()
	server := NewTaskServer()

	router.GET("/task/", server.getAllTasksHandler)

	router.Run("localhost:" + os.Getenv("SERVERPORT"))
}
