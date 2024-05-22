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
)

type MaterialRepo struct {
	MaterialRepo repositories.MaterialFields
}

func NewMaterialFields(repo repositories.MaterialFields) *MaterialRepo {
	return &MaterialRepo{MaterialRepo: repo}
}

func (svc MaterialRepo) CreateMaterialField(e echo.Context) error {
	var req requests.MaterialFields
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusUnprocessableEntity, handlers.ErrorResponse{
			HTTPStatus: http.StatusUnprocessableEntity,
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}
	if req.Name == "" {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    constants.NAME_IS_EMPTY,
		})
	}
	mapper := dao.MaterialFields{
		Name: req.Name,
	}
	exist := svc.MaterialRepo.CheckExisting(mapper)
	if exist == true {
		return e.JSON(http.StatusUnprocessableEntity, handlers.ErrorResponse{
			HTTPStatus: http.StatusUnprocessableEntity,
			Time:       constants.TIME_NOW,
			Message:    constants.MATERIAL_IS_EXIST,
		})
	}
	err := svc.MaterialRepo.Insert(mapper)
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

func (svc MaterialRepo) GetMaterialAll(e echo.Context) error {
	response := svc.MaterialRepo.GetMaterialAll()
	return e.JSON(http.StatusOK, handlers.SuccessResponseListMaterialFields{
		HTTPStatus: http.StatusOK,
		Time:       constants.TIME_NOW,
		Data:       response,
	})
}

func (svc MaterialRepo) GetMaterialById(e echo.Context) error {
	idStr := e.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}
	idUint := uint(idInt)
	response := svc.MaterialRepo.GetMaterialById(idUint)
	return e.JSON(http.StatusOK, handlers.SuccessResponseMaterialFields{
		HTTPStatus: http.StatusOK,
		Time:       constants.TIME_NOW,
		Data:       response,
	})
}
