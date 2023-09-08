package task

import (
	"github.com/emejotaw/go-todo/domain/task/aggregate"
	"github.com/emejotaw/go-todo/domain/task/entity"
	"github.com/emejotaw/go-todo/domain/task/repository"
	"github.com/emejotaw/go-todo/domain/task/repository/memory"
	"github.com/emejotaw/go-todo/valueobject"
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

func (ts *TaskService) CreateTask(dto valueobject.TaskDTO) error {

	todo := aggregate.NewTodo(dto.Name, dto.Description)
	return ts.repository.Create(todo)
}

func (ts *TaskService) GetTasks() ([]entity.Task, error) {

	todos, err := ts.repository.GetAll()
	tasks := []entity.Task{}

	if err != nil {
		return nil, err
	}

	for _, todo := range todos {
		tasks = append(tasks, *todo.GetTask())
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
