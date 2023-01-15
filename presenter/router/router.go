package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/a2-ito/todo-backend/presenter/handler"
)

func SetRouter(e *echo.Echo, h handler.AppHandler) {

	//userController := controllers.NewUserController(NewSqlHandler())
	//userController := controllers.NewUserController()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	r := e.Group("/api")
	r.GET("", h.Hello)
	r.GET("/users", h.GetUsers)
	r.GET("/users/:id", h.GetUser)
	r.POST("/users", h.CreateUser)

	r.GET("/tasks", h.GetTasks)
	r.GET("/tasks/:id", h.GetTask)
	r.POST("/tasks", h.CreateTask)
	r.PUT("/tasks", h.UpdateTask)
	r.DELETE("/tasks/:id", h.DeleteTask)

}
