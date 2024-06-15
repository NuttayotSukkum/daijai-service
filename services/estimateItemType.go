package services

import (
	"daijai-service/constants"
	"daijai-service/models/dao"
	"daijai-service/models/handlers"
	"daijai-service/models/requests"
	"daijai-service/repositories"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strings"
)

type EstimateItemType struct {
	EstimateItemType repositories.EstimateItemType
}

func NewEstimateItemType(repo repositories.EstimateItemType) *EstimateItemType {
	return &EstimateItemType{EstimateItemType: repo}
}

func (svc EstimateItemType) CreateEstimateItemType(e echo.Context) error {
	var req requests.RequestEstimateIItemType

	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    "invalid request body",
		})
	}

	if req.Name == nil || *req.Name == "" {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    "Name is empty",
		})
	}

	materialNameType := svc.EstimateItemType.FindEstimateItemTypesByName(*req.Name)
	log.Printf("Estimate item is not found: %v", req.Name)

	if len(materialNameType.Name) != 0 {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    "EstimateItemType is already exist",
		})
	}

	estimateItemTypeMapper := dao.EstimateItemTypes{
		Name: strings.TrimSpace(*req.Name),
	}

	err := svc.EstimateItemType.Insert(estimateItemTypeMapper)
	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, handlers.ErrorResponse{
			HTTPStatus: http.StatusUnprocessableEntity,
			Time:       constants.TIME_NOW,
			Message:    "create is not success",
		})
	}

	return e.JSON(http.StatusOK, handlers.SuccessReposeMessage{
		HTTPStatus: http.StatusCreated,
		Time:       constants.TIME_NOW,
		Name:       *req.Name,
		Message:    "success",
	})
}

func (svc EstimateItemType) GetEstimateItemTypeAll(e echo.Context) error {
	estimateTypeAll, err := svc.EstimateItemType.FindEstimateItemTypeAll()
	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, handlers.ErrorResponse{
			HTTPStatus: http.StatusUnprocessableEntity,
			Time:       constants.TIME_NOW,
			Message:    "EstimateItemType is empty",
		})
	}

	return e.JSON(http.StatusOK, handlers.SuccessResponseEstimateItemTypeList{
		HTTPStatus: http.StatusOK,
		Time:       constants.TIME_NOW,
		Data:       estimateTypeAll,
	})

}
