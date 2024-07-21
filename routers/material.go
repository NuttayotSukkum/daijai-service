package routers

import (
	"daijai-service/configs"
	"daijai-service/middleware"
	"daijai-service/repositories/db"
	"daijai-service/services"
	"github.com/labstack/echo/v4"
)

func MaterialRouter(e *echo.Echo) {
	configs.GetDBInstance()
	dbInstance := configs.GetDBInstance()

	materialRepo := db.MaterialRepository(dbInstance)
	cat3Repo := db.NewCategory3Repository(dbInstance)
	materialDetailsRepo := db.NewMaterialDetails(dbInstance)
	materialFieldDetailRepo := db.NewMaterialDetailsRepository(dbInstance)
	materialSvc := services.NewMaterials(materialRepo, cat3Repo, materialDetailsRepo, materialFieldDetailRepo)

	g := e.Group("/user", middleware.ValidateTokenMiddleware)
	g.POST("/v1/daijai/material/create", materialSvc.CreateMaterials)
	g.GET("/v1/daijai/material/materials", materialSvc.GetMaterials)
}
