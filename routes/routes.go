package routes

import (
	"github.com/emejotaw/go-todo/api"
	"github.com/gin-gonic/gin"
)

func Start() {

	taskController := api.NewTaskController()

	server := gin.Default()
	server.GET("/tasks", taskController.GetTasks)

	server.Run(":3000")
}
