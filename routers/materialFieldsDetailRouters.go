package routers

import (
	"daijai-service/configs"
	"daijai-service/middleware"
	"daijai-service/repositories/db"
	"daijai-service/services"
	"github.com/labstack/echo/v4"
)

func MaterialDetailRouter(e *echo.Echo) {
	configs.GetDBInstance()
	dbInstance := configs.GetDBInstance()

	materialDetailRepo := db.NewMaterialDetailsRepository(dbInstance)
	materialSvc := services.NewMaterialDetailRepository(materialDetailRepo)

	g := e.Group("/user", middleware.ValidateTokenMiddleware)
	g.POST("/v1/daijai/material_field_details/create", materialSvc.CreateMaterialDetail)
	g.GET("/v1/daijai/material_field_details", materialSvc.GetListMaterial)
	g.GET("/v1/daijai/material_field_details/:id", materialSvc.GetMaterial)
	g.GET("/v1/daijai//material_field_details/material_field/:id", materialSvc.GetMaterialByMaterialFields)
}
