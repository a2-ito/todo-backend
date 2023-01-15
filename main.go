package main

import (
	_ "github.com/a2-ito/todo-backend/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/a2-ito/todo-backend/interactor"
	"github.com/a2-ito/todo-backend/presenter/router"
)

// @title          Todo API
// @version        0.1
// @description    Todo 内部 API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host           localhost:8080
// @BasePath       /
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
