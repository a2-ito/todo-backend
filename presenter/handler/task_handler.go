package handler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/a2-ito/todo-backend/domain/model"
	"github.com/a2-ito/todo-backend/usecase"
)

type TaskHandler interface {
	GetTasks(c echo.Context) error
	GetTask(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskHandler struct {
	TaskUseCase usecase.TaskUseCase
}

func NewTaskHandler(u usecase.TaskUseCase) TaskHandler {
	return &taskHandler{u}
}

func (h *taskHandler) GetTask(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))
	fmt.Println("handler GetTask ", id)

	task, err := h.TaskUseCase.GetTask(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)
}

func (h *taskHandler) GetTasks(c echo.Context) error {
	userId := uuid.MustParse(c.Param("id"))

	task, err := h.TaskUseCase.GetTasks(userId)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)
}

type (
	Task struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}
)

func (h *taskHandler) CreateTask(c echo.Context) error {
	task := &model.Task{}
	if err := c.Bind(task); err != nil {
		return err
	}

	createdTask, err := h.TaskUseCase.CreateTask(task)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, createdTask)
}

func (h *taskHandler) UpdateTask(c echo.Context) error {
	task := &model.Task{}
	if err := c.Bind(task); err != nil {
		return err
	}

	updatedTask, err := h.TaskUseCase.UpdateTask(task)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func (h *taskHandler) DeleteTask(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))

	err := h.TaskUseCase.DeleteTask(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "delete task")
}
