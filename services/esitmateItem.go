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

type EstimateItem struct {
	EstimateItemType repositories.EstimateItemType
	EstimateItem     repositories.EstimateItem
}

func NewEstimateItem(EstimateItemRepo repositories.EstimateItem, estimateItemTypeRepo repositories.EstimateItemType) *EstimateItem {
	return &EstimateItem{
		EstimateItem:     EstimateItemRepo,
		EstimateItemType: estimateItemTypeRepo,
	}
}

func (svc *EstimateItem) CreateEstimateItem(e echo.Context) error {
	var req requests.RequestEstimateItem
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}

	if req.EstimateItemTypeId == 0 {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    "EstimateItemTypeId is required",
		})
	}

	if req.Name == nil || *req.Name == "" {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    "Name is required",
		})
	}

	if req.Code == nil || len(*req.Code) == 0 {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    "Code is required",
		})
	}

	estimateItemType, err := svc.EstimateItemType.FindEstimateItemTypesById(req.EstimateItemTypeId)
	if err != nil {
		log.Printf("EstimateItemTypeId:%d error:%v", req.EstimateItemTypeId, err)
	}
	if len(estimateItemType) == 0 {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    "EstimateItemType doesn't exists",
		})
	}

	estimateItemMapper := dao.EstimateItem{
		EstimateItemTypeId: req.EstimateItemTypeId,
		Name:               strings.TrimSpace(*req.Name),
		Code:               strings.TrimSpace(strings.ToUpper(*req.Code)),
		Price:              req.Price,
	}

	err = svc.EstimateItem.Insert(estimateItemMapper)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, handlers.ErrorResponse{
			HTTPStatus: http.StatusInternalServerError,
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}

	return e.JSON(
		http.StatusOK, handlers.SuccessReposeMessage{
			HTTPStatus: http.StatusOK,
			Time:       constants.TIME_NOW,
			Name:       *req.Name,
			Message:    "EstimateItem created successfully",
		})

}

func (svc *EstimateItem) GetAll(e echo.Context) error {
	estimateItemAll, err := svc.EstimateItem.FindEstimateItemAll()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, handlers.ErrorResponse{
			HTTPStatus: http.StatusInternalServerError,
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}
	return e.JSON(http.StatusOK, handlers.SuccessResponseEstimateItemList{
		HTTPStatus: http.StatusOK,
		Time:       constants.TIME_NOW,
		Data:       estimateItemAll,
	})
}
