package aggregate

import (
	"github.com/emejotaw/go-todo/domain/task/entity"
	"github.com/google/uuid"
)

type TODO struct {
	task *entity.Task
}

func (t *TODO) GetID() uuid.UUID {
	return t.task.ID
}
