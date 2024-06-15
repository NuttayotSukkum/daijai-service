package routers

import (
	"daijai-service/configs"
	"daijai-service/middleware"
	"daijai-service/repositories/db"
	"daijai-service/services"
	"github.com/labstack/echo/v4"
)

func EstimateItem(e *echo.Echo) {
	configs.GetDBInstance()
	dbInstance := configs.GetDBInstance()

	estimateItemRepo := db.NewEstimateItem(dbInstance)
	estimateItemTypeRepo := db.NewEstimateItemType(dbInstance)
	estimateItemSvc := services.NewEstimateItem(estimateItemRepo, estimateItemTypeRepo)

	g := e.Group("/user", middleware.ValidateTokenMiddleware)
	g.POST("/v1/daijai/estimate_item/create", estimateItemSvc.CreateEstimateItem)
	g.GET("/v1/daijai/estimate_item/estimateitems", estimateItemSvc.GetAll)
}
