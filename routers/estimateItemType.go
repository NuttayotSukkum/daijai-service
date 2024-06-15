package routers

import (
	"daijai-service/configs"
	"daijai-service/middleware"
	"daijai-service/repositories/db"
	"daijai-service/services"
	"github.com/labstack/echo/v4"
)

func EstimateItemType(e *echo.Echo) {
	configs.GetDBInstance()
	dbInstance := configs.GetDBInstance()

	estimateItemTypeRepo := db.NewEstimateItemType(dbInstance)
	estimateItemTypeSvc := services.NewEstimateItemType(estimateItemTypeRepo)

	g := e.Group("/user", middleware.ValidateTokenMiddleware)
	g.POST("/v1/daijai/estimate_item_type/create", estimateItemTypeSvc.CreateEstimateItemType)
	g.GET("/v1/daijai/estimate_item_type/estimateitemtypes", estimateItemTypeSvc.GetEstimateItemTypeAll)
}
