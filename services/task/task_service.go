package task

import (
	"github.com/emejotaw/go-todo/domain/task/repository"
	"github.com/emejotaw/go-todo/domain/task/repository/memory"
)

type TaskConfiguration func(ts *TaskService) error

type TaskService struct {
	repository repository.Repository
}

func (ts *TaskService) CreateTask() {

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
