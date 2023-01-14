package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/a2-ito/todo-backend/interactor"
	"github.com/a2-ito/todo-backend/presenter/router"
)

func main() {
	e := echo.New()

	e.Logger.SetLevel(log.DEBUG)

	/*
		middleware.NewMiddleware(e)

		conn := config.NewDBConnection(cfg)
		conf := config.NewOauthConfig(cfg)
		key := cfg.SigningKey
	*/

	i := interactor.NewInteractor()
	h := i.NewAppHandler()

	router.SetRouter(e, h)

	e.Logger.Fatal(e.Start(":1323"))

}
