package repository

import (
	"github.com/a2-ito/todo-backend/domain/model"
	"github.com/google/uuid"
)

type TaskRepository interface {
	FindById(id uuid.UUID) (*model.Task, error)
	FindAllByUserId(id uuid.UUID) ([]*model.Task, error)
	Create(*model.Task) (*model.Task, error)
	Update(*model.Task) (*model.Task, error)
	Delete(id uuid.UUID) error
}
