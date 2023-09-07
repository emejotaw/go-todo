package aggregate

import (
	"errors"

	"github.com/emejotaw/go-todo/domain/task/entity"
	"github.com/google/uuid"
)

var (
	ErrInvalidInputData = errors.New("invalid input data")
)

type TODO struct {
	task *entity.Task
}

func NewTodo(name, description string) *TODO {

	task := &entity.Task{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
	}

	return &TODO{
		task: task,
	}
}

func NewTodoWithTask(task *entity.Task) *TODO {

	return &TODO{
		task: task,
	}
}

func (t *TODO) GetID() uuid.UUID {
	return t.task.ID
}
