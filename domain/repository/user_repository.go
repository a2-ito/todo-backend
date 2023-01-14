package repository

import (
	"github.com/a2-ito/todo-backend/domain/model"
	"github.com/google/uuid"
)

type UserRepository interface {
	Hello()
	FindById(id uuid.UUID) (*model.User, error)
	FindAll() ([]*model.User, error)
	Create(*model.User) (*model.User, error)
	/*
	  Update(domain.User) (*domain.User, error)
	  Delete(id uuid.UUID) error
	*/
}
