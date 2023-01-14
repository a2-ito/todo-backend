package interactor

import (
	"github.com/a2-ito/todo-backend/domain/repository"
	"github.com/a2-ito/todo-backend/infrastructure/database/memory"
	"github.com/a2-ito/todo-backend/presenter/handler"
	"github.com/a2-ito/todo-backend/usecase"
)

type Interactor interface {
	NewAppHandler() handler.AppHandler
	NewUserHandler() handler.UserHandler
	NewUserUseCase() usecase.UserUseCase
	NewTaskHandler() handler.TaskHandler
	NewTaskUseCase() usecase.TaskUseCase
}

type interactor struct {
}

func NewInteractor() Interactor {
	return &interactor{}
}

type appHandler struct {
	handler.UserHandler
	handler.TaskHandler
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.UserHandler = i.NewUserHandler()
	appHandler.TaskHandler = i.NewTaskHandler()

	return appHandler
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserUseCase())
}

func (i *interactor) NewTaskHandler() handler.TaskHandler {
	return handler.NewTaskHandler(i.NewTaskUseCase())
}

func (i *interactor) NewUserRepository() repository.UserRepository {
	return memory.NewUserRepository()
}

func (i *interactor) NewUserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(i.NewUserRepository())
}

func (i *interactor) NewTaskRepository() repository.TaskRepository {
	return memory.NewTaskRepository()
}

func (i *interactor) NewTaskUseCase() usecase.TaskUseCase {
	return usecase.NewTaskUseCase(i.NewTaskRepository())
}
