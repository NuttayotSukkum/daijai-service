package routers

import (
	"daijai-service/configs"
	"daijai-service/middleware"
	"daijai-service/repositories/db"
	"daijai-service/services"
	"github.com/labstack/echo/v4"
)

func MaterialFieldsRouter(e *echo.Echo) {
	configs.GetDBInstance()
	dbInstance := configs.GetDBInstance()

	materialRepo := db.NewMaterialFieldsRepository(dbInstance)
	materialSvc := services.NewMaterialFields(materialRepo)

	g := e.Group("/user", middleware.ValidateTokenMiddleware)
	g.POST("/v1/daijai/material_fields/create", materialSvc.CreateMaterialField)
	g.GET("/v1/daijai/material_fields/material_fields", materialSvc.GetMaterialAll)
	g.GET("/v1/daijai/material_fields/material_fields/:id", materialSvc.GetMaterialById)

}
