package routers

import (
	"daijai-service/configs"
	"daijai-service/middleware"
	"daijai-service/services"
	"github.com/labstack/echo/v4"
)

func ProjectRouter() *echo.Echo {
	e := echo.New()
	configs.GetDBInstance()

	g := e.Group("/user", middleware.ValidateTokenMiddleware)
	g.POST("/v1/daijai/project", services.CeateProject)
	return e
}

func Execute(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":8080"))
}
