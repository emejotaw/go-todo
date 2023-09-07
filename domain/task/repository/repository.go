package repository

import (
	"errors"

	"github.com/emejotaw/go-todo/domain/task/aggregate"
)

var (
	ErrTaskAlreadyExists = errors.New("task already exists")
)

type Repository interface {
	Create(aggregate.TODO) error
	GetAll() (aggregate.TODO, error)
}
