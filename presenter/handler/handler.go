package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/a2-ito/todo-backend/domain/model"
	"github.com/a2-ito/todo-backend/usecase"
)

type UserHandler interface {
	Hello(c echo.Context) error
	GetUsers(c echo.Context) error
	GetUser(c echo.Context) error
	CreateUser(c echo.Context) error
}

type userHandler struct {
	UserUseCase usecase.UserUseCase
}

func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return &userHandler{u}
}

func (h *userHandler) Hello(c echo.Context) error {
	fmt.Println("handler Hello")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	h.UserUseCase.Hello(ctx)
	return c.JSON(http.StatusOK, "hello")
}

func (h *userHandler) GetUser(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))
	fmt.Println("handler GetUser ", id)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.UserUseCase.GetUser(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) GetUsers(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.UserUseCase.GetUsers(ctx)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

type (
	User struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}
)

func (h *userHandler) CreateUser(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return err
	}

	createdUser, err := h.UserUseCase.CreateUser(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, createdUser)
}
