package routers

import (
	"daijai-service/configs"
	"daijai-service/middleware"
	"daijai-service/repositories/db"
	"daijai-service/services"
	"github.com/labstack/echo/v4"
)

func Category3Router(e *echo.Echo) {
	configs.GetDBInstance()
	dbInstance := configs.GetDBInstance()

	category3Repo := db.NewCategory3Repository(dbInstance)
	category3Svc := services.NewCategory3Service(category3Repo)

	g := e.Group("/user", middleware.ValidateTokenMiddleware)
	g.POST("/v1/daijai/category_3s/create", category3Svc.CreateCategory3)
	g.GET("/v1/daijai/category_3s/category_3s", category3Svc.GetCategory3All)
	g.GET("/v1/daijai/category_3s/:id", category3Svc.GetCat3)

}
