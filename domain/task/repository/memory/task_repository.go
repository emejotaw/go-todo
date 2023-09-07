package memory

import (
	"github.com/emejotaw/go-todo/domain/task/aggregate"
	"github.com/emejotaw/go-todo/domain/task/repository"
	"github.com/google/uuid"
)

type TaskRepository struct {
	tasks map[uuid.UUID]aggregate.TODO
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[uuid.UUID]aggregate.TODO),
	}
}

func (t *TaskRepository) Create(todo aggregate.TODO) error {

	if t.tasks == nil {
		t.tasks = make(map[uuid.UUID]aggregate.TODO)
	}

	if _, ok := t.tasks[todo.GetID()]; ok {
		return repository.ErrTaskAlreadyExists
	}

	t.tasks[todo.GetID()] = todo
	return nil
}

func (t *TaskRepository) GetAll() ([]aggregate.TODO, error) {

	tasks := []aggregate.TODO{}

	for _, task := range t.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}
