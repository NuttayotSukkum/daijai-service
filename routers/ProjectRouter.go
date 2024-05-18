package routers

import (
	"daijai-service/configs"
	"daijai-service/middleware"
	"daijai-service/repositories/db"
	"daijai-service/services"
	"github.com/labstack/echo/v4"
)

func ProjectRouter() *echo.Echo {
	e := echo.New()

	configs.GetDBInstance()
	dbInstance := configs.GetDBInstance()

	projectStatusRepo := db.NewProjectStatusRepository(dbInstance)
	projectStatusSvc := services.NewProjectStatusService(projectStatusRepo)

	g := e.Group("/user", middleware.ValidateTokenMiddleware)
	g.POST("/v1/daijai/project", projectStatusSvc.CreateProject)
	g.GET("/v1/daijai/:projectId", projectStatusSvc.GetProjectStatus)
	g.GET("/v1/daijai/getall", projectStatusSvc.GetAllProject)
	return e
}

func Execute(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":8080"))
}
