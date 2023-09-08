package routes

import (
	"github.com/emejotaw/go-todo/api"
	"github.com/emejotaw/go-todo/services/task"
	"github.com/gin-gonic/gin"
)

func Start() {

	taskController := api.NewTaskController(
		task.NewTaskService(
			task.WithMemoryRepository()),
	)

	server := gin.Default()
	server.GET("/tasks", taskController.GetTasks)
	server.POST("/tasks", taskController.CreateTask)

	server.Run(":3000")
}
