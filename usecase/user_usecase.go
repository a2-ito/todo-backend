package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	// "github.com/labstack/echo/v4"
	"github.com/a2-ito/todo-backend/domain/model"
	"github.com/a2-ito/todo-backend/domain/repository"
)

/*
type UserRepository interface {
  Hello() error
  //FindById(id string) (domain.User, error)
}

type UserInteractor struct {
  UserRepository repository.UserRepository
}
*/

type UserUseCase interface {
	Hello(ctx context.Context)
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetUser(ctx context.Context, id uuid.UUID) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
}

type userUseCase struct {
	repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{r}
}

func (usecase *userUseCase) Hello(ctx context.Context) {
	fmt.Println("UseCase Hello")
	usecase.UserRepository.Hello()
}

func (usecase *userUseCase) GetUser(ctx context.Context, id uuid.UUID) (*model.User, error) {
	fmt.Println("UseCase FindById ", id)
	return usecase.UserRepository.FindById(id)
}

func (usecase *userUseCase) GetUsers(ctx context.Context) ([]*model.User, error) {
	fmt.Println("UseCase FindAll")
	return usecase.UserRepository.FindAll()
}

func (usecase *userUseCase) CreateUser(user *model.User) (*model.User, error) {
	return usecase.UserRepository.Create(user)
}
