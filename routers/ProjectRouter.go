package routers

import (
	"daijai-service/configs"
	"daijai-service/middleware"
	"daijai-service/repositories/db"
	"daijai-service/services"
	"github.com/labstack/echo/v4"
)

func ProjectRouter(e *echo.Echo) {

	configs.GetDBInstance()
	dbInstance := configs.GetDBInstance()

	projectStatusRepo := db.NewProjectStatusRepository(dbInstance)
	projectStatusSvc := services.NewProjectStatusService(projectStatusRepo)

	g := e.Group("/user", middleware.ValidateTokenMiddleware)
	g.POST("/v1/daijai/projects/create", projectStatusSvc.CreateProject)
	g.GET("/v1/daijai/project/:id", projectStatusSvc.GetProjectStatus)
	g.GET("/v1/daijai/projects", projectStatusSvc.GetAllProject)
	g.PUT("/v1/daijai/projects/update", projectStatusSvc.UpdateProject)
	//g.DELETE("/v1/daijai/project/delete-project", projectStatusSvc.DeleteProject)

}

func Execute(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":8080"))
}
