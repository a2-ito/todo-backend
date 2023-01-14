package handler

type (
	AppHandler interface {
		UserHandler
		TaskHandler
	}
)
