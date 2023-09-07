package aggregate

import (
	"testing"

	"github.com/emejotaw/go-todo/domain/task/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewTodo(t *testing.T) {

	type testCase struct {
		name        string
		description string
	}

	tests := []testCase{
		{name: "Task 1", description: "Task 1"},
		{name: "Task 2", description: "Task 2"},
		{name: "Task 3", description: "Task 3"},
		{name: "Task 4", description: "Task 4"},
		{name: "Task 5", description: "Task 5"},
	}

	for _, test := range tests {

		todo := NewTodo(test.name, test.description)
		assert.NotNil(t, todo)
	}
}

func TestNewTodoWithTask(t *testing.T) {

	type testCase struct {
		name        string
		description string
	}

	tests := []testCase{
		{name: "Task 1", description: "Task 1"},
		{name: "Task 2", description: "Task 2"},
		{name: "Task 3", description: "Task 3"},
		{name: "Task 4", description: "Task 4"},
		{name: "Task 5", description: "Task 5"},
	}

	for _, test := range tests {

		task := &entity.Task{
			ID:          uuid.New(),
			Name:        test.name,
			Description: test.description,
		}
		todo := NewTodoWithTask(task)
		assert.NotNil(t, todo)
	}
}

func TestGetID(t *testing.T) {

	task := &entity.Task{
		ID:          uuid.New(),
		Name:        "Some name",
		Description: "Some description",
	}

	todo := TODO{task: task}
	assert.NotEmpty(t, todo.GetID())
}
