package main

import (
	"daijai-service/routers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routers.ProjectRouter(e)
	routers.Category3Router(e)
	routers.MaterialFieldsRouter(e)
	routers.MaterialDetailRouter(e)
	routers.MaterialRouter(e)

	routers.Execute(e)
}
