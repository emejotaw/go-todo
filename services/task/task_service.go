package task

import (
	"github.com/emejotaw/go-todo/domain/task/aggregate"
	"github.com/emejotaw/go-todo/domain/task/repository"
	"github.com/emejotaw/go-todo/domain/task/repository/memory"
)

type TaskConfiguration func(ts *TaskService) error

type TaskService struct {
	repository repository.Repository
}

func NewTaskService(configurations ...TaskConfiguration) *TaskService {

	taskService := &TaskService{}

	for _, configuration := range configurations {
		configuration(taskService)
	}

	return taskService
}

func (ts *TaskService) CreateTask() {

}

func (ts *TaskService) GetTasks() ([]aggregate.TODO, error) {

	tasks, err := ts.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return tasks, err
}

func WithRepository(repository repository.Repository) TaskConfiguration {

	return func(ts *TaskService) error {
		ts.repository = repository
		return nil
	}
}

func WithMemoryRepository() TaskConfiguration {

	repository := memory.NewTaskRepository()
	return WithRepository(repository)
}
