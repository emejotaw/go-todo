package api

import (
	"net/http"

	"github.com/emejotaw/go-todo/services/task"
	"github.com/emejotaw/go-todo/valueobject"
	"github.com/gin-gonic/gin"
)

type taskController struct {
	taskService *task.TaskService
}

func NewTaskController(taskService *task.TaskService) *taskController {
	return &taskController{taskService: taskService}
}

func (t *taskController) CreateTask(c *gin.Context) {

	dto := &valueobject.TaskDTO{}
	c.BindJSON(dto)
	taskService := t.taskService
	err := taskService.CreateTask(*dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    999,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    0,
		"message": "Task created successfully",
	})
}

func (t *taskController) GetTasks(c *gin.Context) {

	taskService := t.taskService
	tasks, err := taskService.GetTasks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}
