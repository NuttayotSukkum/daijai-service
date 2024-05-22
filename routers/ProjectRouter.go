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
	g.GET("/v1/daijai/:projectName", projectStatusSvc.GetProjectStatus)
	g.GET("/v1/daijai/getall", projectStatusSvc.GetAllProject)
	g.PUT("/v1/daijai/update-project", projectStatusSvc.UpdateProject)
	//g.DELETE("/v1/daijai/project/delete-project", projectStatusSvc.DeleteProject)
	return e
}

func Execute(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":8080"))
}
