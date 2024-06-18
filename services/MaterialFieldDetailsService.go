package services

import (
	"daijai-service/constants"
	"daijai-service/models/dao"
	"daijai-service/models/handlers"
	"daijai-service/models/requests"
	"daijai-service/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

type MaterialDetailsRepo struct {
	MaterialDetailsRepo repositories.MaterialFieldsDetail
}

func NewMaterialDetailRepository(repo repositories.MaterialFieldsDetail) *MaterialDetailsRepo {
	return &MaterialDetailsRepo{MaterialDetailsRepo: repo}
}

func (svc MaterialDetailsRepo) CreateMaterialDetail(e echo.Context) error {
	var req requests.RequestMaterialDetail
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusUnprocessableEntity, handlers.ErrorResponse{
			HTTPStatus: http.StatusUnprocessableEntity,
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}
	if req.MaterialFieldsId == "" {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    constants.MATERIAL_FIELDS_IS_EMPTY,
		})
	}
	if req.Name == "" {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    constants.NAME_IS_EMPTY,
		})
	}
	if req.Code == "" {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    constants.CODE_EMPTY,
		})
	}

	materialFieldIdInt, err := strconv.Atoi(req.MaterialFieldsId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    constants.CAN_NOT_CONVERT,
		})
	}

	mapper := dao.MaterialFieldDetail{
		MaterialFieldId: materialFieldIdInt,
		Name:            strings.TrimSpace(strings.ToUpper(req.Name)),
		Code:            strings.TrimSpace(req.Code),
	}
	exist := svc.MaterialDetailsRepo.CheckExisting(req.Name)

	if exist == true {
		return e.JSON(http.StatusUnprocessableEntity, handlers.ErrorResponse{
			HTTPStatus: http.StatusUnprocessableEntity,
			Time:       constants.TIME_NOW,
			Message:    constants.MATERIAL_IS_EXIST,
		})
	}

	err = svc.MaterialDetailsRepo.Insert(mapper)
	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, handlers.ErrorResponse{
			HTTPStatus: http.StatusUnprocessableEntity,
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}
	return e.JSON(http.StatusOK, handlers.ErrorResponse{
		HTTPStatus: http.StatusOK,
		Time:       constants.TIME_NOW,
		Message:    constants.CREATE_SUCCESS,
	})
}

func (svc MaterialDetailsRepo) GetListMaterial(e echo.Context) error {
	response := svc.MaterialDetailsRepo.GetMaterialAll()
	return e.JSON(http.StatusOK, handlers.SuccessResponseMaterialFieldDetailsList{
		HTTPStatus: http.StatusOK,
		Time:       constants.TIME_NOW,
		Data:       response,
	})

}

func (svc MaterialDetailsRepo) GetMaterial(e echo.Context) error {
	idStr := e.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, handlers.ErrorResponse{
			HTTPStatus: http.StatusUnprocessableEntity,
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}
	response := svc.MaterialDetailsRepo.GetMaterialById(idInt)
	if len(response.Code) == 0 {
		return e.JSON(http.StatusNotFound, handlers.ErrorResponse{
			HTTPStatus: http.StatusNotFound,
			Time:       constants.TIME_NOW,
			Message:    constants.NOT_FOUND_DATA,
		})
	}
	return e.JSON(http.StatusOK, handlers.SuccessResponseMaterialFieldDetails{
		HTTPStatus: http.StatusOK,
		Time:       constants.TIME_NOW,
		Data:       response,
	})
}

func (svc MaterialDetailsRepo) GetMaterialByMaterialFields(e echo.Context) error {
	idStr := e.Param("id")
	response := svc.MaterialDetailsRepo.GetMaterialByMaterialFields(idStr)
	return e.JSON(http.StatusOK, handlers.SuccessResponseMaterialFieldDetails{
		HTTPStatus: http.StatusOK,
		Time:       constants.TIME_NOW,
		Data:       response,
	})
}
