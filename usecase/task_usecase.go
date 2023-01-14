package usecase

import (
	"fmt"

	"github.com/google/uuid"
	// "github.com/labstack/echo/v4"
	"github.com/a2-ito/todo-backend/domain/model"
	"github.com/a2-ito/todo-backend/domain/repository"
)

type TaskUseCase interface {
	GetTasks(id uuid.UUID) ([]*model.Task, error)
	GetTask(id uuid.UUID) (*model.Task, error)
	CreateTask(task *model.Task) (*model.Task, error)
	UpdateTask(task *model.Task) (*model.Task, error)
	DeleteTask(id uuid.UUID) error
}

type taskUseCase struct {
	repository.TaskRepository
}

func NewTaskUseCase(r repository.TaskRepository) TaskUseCase {
	return &taskUseCase{r}
}

func (usecase *taskUseCase) GetTask(id uuid.UUID) (*model.Task, error) {
	fmt.Println("UseCase FindById ", id)
	return usecase.TaskRepository.FindById(id)
}

func (usecase *taskUseCase) GetTasks(id uuid.UUID) ([]*model.Task, error) {
	fmt.Println("UseCase FindAll")

	return usecase.TaskRepository.FindAllByUserId(id)
}

func (usecase *taskUseCase) CreateTask(task *model.Task) (*model.Task, error) {
	return usecase.TaskRepository.Create(task)
}

func (usecase *taskUseCase) UpdateTask(task *model.Task) (*model.Task, error) {
	return usecase.TaskRepository.Update(task)
}

func (usecase *taskUseCase) DeleteTask(id uuid.UUID) error {
	return usecase.TaskRepository.Delete(id)
}
