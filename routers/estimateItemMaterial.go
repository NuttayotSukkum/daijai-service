package routers

import (
	"daijai-service/configs"
	"daijai-service/middleware"
	"daijai-service/repositories/db"
	"daijai-service/services"
	"github.com/labstack/echo/v4"
)

func EstimateItemMaterial(e *echo.Echo) {
	configs.GetDBInstance()
	dbInstance := configs.GetDBInstance()

	estimateItemRepo := db.NewEstimateItem(dbInstance)
	estimateItemMaterialRepo := db.NewEstimateItemMaterial(dbInstance)
	materialRepo := db.MaterialRepository(dbInstance)
	estimateItemMaterialSvc := services.NewEstimateItemMaterials(estimateItemMaterialRepo, materialRepo, estimateItemRepo)

	g := e.Group("/user", middleware.ValidateTokenMiddleware)
	g.POST("/v1/daijai/estimate_item_material/create", estimateItemMaterialSvc.CreateItemMaterial)
	//g.GET("/v1/daijai/estimate_item_material/estimate_item_materials", estimateItemMaterialSvc.GetAll)
}
