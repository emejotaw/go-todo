package api

import (
	"net/http"

	"github.com/emejotaw/go-todo/services/task"
	"github.com/gin-gonic/gin"
)

type taskController struct {
}

func NewTaskController() *taskController {
	return &taskController{}
}

func (t *taskController) GetTasks(c *gin.Context) {

	taskService := task.NewTaskService(task.WithMemoryRepository())
	tasks, err := taskService.GetTasks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}
