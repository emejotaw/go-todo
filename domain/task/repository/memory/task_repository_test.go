package memory

import (
	"testing"

	"github.com/emejotaw/go-todo/domain/task/aggregate"
	"github.com/emejotaw/go-todo/domain/task/entity"
	"github.com/emejotaw/go-todo/domain/task/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {

	repository := NewTaskRepository()
	assert.NotNil(t, repository)
}

func TestCreate(t *testing.T) {

	type testCase struct {
		ID          uuid.UUID
		Name        string
		Description string
		ExpectedErr error
	}

	duplicateUUID, _ := uuid.Parse("5492f388-372c-4016-9517-c89c23cbc07d")
	tests := []testCase{
		{ID: duplicateUUID, Name: "Task 1", Description: "Task 1", ExpectedErr: nil},
		{ID: uuid.New(), Name: "Task 2", Description: "Task 2", ExpectedErr: nil},
		{ID: uuid.New(), Name: "Task 3", Description: "Task 3", ExpectedErr: nil},
		{ID: uuid.New(), Name: "Task 4", Description: "Task 4", ExpectedErr: nil},
		{ID: duplicateUUID, Name: "Task 4", Description: "Task 4", ExpectedErr: repository.ErrTaskAlreadyExists},
	}

	repository := TaskRepository{}
	for _, test := range tests {

		task := &entity.Task{
			ID:          test.ID,
			Name:        test.Name,
			Description: test.Description,
		}
		todo := aggregate.NewTodoWithTask(task)
		err := repository.Create(todo)

		assert.Equal(t, test.ExpectedErr, err)
	}
}

func TestGetAll(t *testing.T) {

	type testCase struct {
		Name        string
		Description string
		ExpectedErr error
	}

	tests := []testCase{
		{Name: "Task 1", Description: "Task 1", ExpectedErr: nil},
		{Name: "Task 2", Description: "Task 2", ExpectedErr: nil},
		{Name: "Task 3", Description: "Task 3", ExpectedErr: nil},
		{Name: "Task 4", Description: "Task 4", ExpectedErr: nil},
	}

	repository := NewTaskRepository()

	for _, test := range tests {
		todo := aggregate.NewTodo(test.Name, test.Description)
		repository.Create(todo)
	}

	tasks, err := repository.GetAll()

	if err != nil {
		t.Errorf("Error retrieving data, error %v", err.Error())
	}

	assert.Len(t, tasks, 4)
}
