package services

import (
	"daijai-service/constants"
	"daijai-service/models/dao"
	"daijai-service/models/handlers"
	"daijai-service/models/requests"
	"daijai-service/repositories"
	"daijai-service/utils"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type EstimateItemMaterials struct {
	Project               repositories.ProjectStatus
	EstimateItem          repositories.EstimateItem
	EstimateItemMaterials repositories.EstimateItemMaterial
	Material              repositories.Material
}

func NewEstimateItemMaterials(estimateItemMaterialRepo repositories.EstimateItemMaterial, materialsRepo repositories.Material, estimateItems repositories.EstimateItem, projectRepo repositories.ProjectStatus) *EstimateItemMaterials {
	return &EstimateItemMaterials{
		Project:               projectRepo,
		EstimateItem:          estimateItems,
		EstimateItemMaterials: estimateItemMaterialRepo,
		Material:              materialsRepo,
	}
}

func (svc *EstimateItemMaterials) CreateItemMaterial(e echo.Context) error {
	var req requests.RequestEstimateItemMaterial
	if err := e.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}
	if req.ProjectId == 0 {
		return e.JSON(http.StatusNotFound, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    "Project Id is empty",
		})
	}
	if req.EstimateItemId == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    "materialId can not be null",
		})
	}

	if req.MaterialId == nil {
		return echo.NewHTTPError(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    "materialId can not be null",
		})
	}
	project, err := svc.Project.GetByProjectId(req.ProjectId)
	if err != nil {
		log.Println("err: %v", err)
	}

	if &project.Id == nil {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    "Project Id does not exist",
		})
	}

	estimateItemExist := svc.EstimateItem.FindEstimateItemExist(req.EstimateItemId)

	if estimateItemExist == nil {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    "Estimate Item is not exist",
		})
	}
	estimateItemMaterialMapper := dao.EstimateItemMaterial{
		ProjectId:      req.ProjectId,
		EstimateItemId: req.EstimateItemId,
		MaterialAmount: req.MaterialAmount,
		MaterialUnit:   req.MaterialUnit,
	}

	var materials []dao.Material
	var estimateItem dao.EstimateItemMaterial

	for _, code := range req.MaterialId {
		materialExist := svc.Material.FindMaterialsExist(code)

		if materialExist == nil {
			return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
				HTTPStatus: http.StatusBadRequest,
				Time:       constants.TIME_NOW,
				Message:    "Material is not exist",
			})
		}

		estimateItemMaterialMapper.MaterialId = code
		estimateItemMaterial, err := svc.EstimateItemMaterials.Insert(estimateItemMaterialMapper)
		estimateItem = estimateItemMaterial
		if err != nil {
			return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
				HTTPStatus: http.StatusBadRequest,
				Time:       constants.TIME_NOW,
				Message:    err.Error(),
			})
		}

		materialAll := svc.Material.FindMaterialsById(code)
		materials = append(materials, materialAll)

		log.Println("estimateItem", estimateItem)
	}
	projectDetail := utils.ObjectMapper(estimateItem, materials)
	log.Println(projectDetail)
	return e.JSON(http.StatusOK, handlers.SuccessResponseEstimateItemMaterial{
		HTTPStatus: http.StatusOK,
		Time:       constants.TIME_NOW,
		Data:       projectDetail,
	})
}
