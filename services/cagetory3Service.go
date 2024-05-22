package services

import (
	"daijai-service/constants"
	"daijai-service/models/dao"
	"daijai-service/models/handlers"
	"daijai-service/models/requests"
	"daijai-service/repositories"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type Category3Repo struct {
	Category3Repo repositories.Category3
}

func NewCategory3Service(repo repositories.Category3) *Category3Repo {
	return &Category3Repo{Category3Repo: repo}
}

func (svc Category3Repo) CreateCategory3(e echo.Context) error {
	var rq requests.RequestCategory3

	if err := e.Bind(&rq); err != nil {
		return e.JSON(http.StatusUnprocessableEntity, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}

	if rq.Name == "" {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    constants.NAME_IS_EMPTY,
		})
	}
	if rq.Code == "" {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    constants.CODE_EMPTY,
		})
	}
	cat3 := dao.Category3{
		Name: rq.Name,
		Code: rq.Code,
	}
	log.Println("Cate3:{}", cat3)
	cat3Existing := svc.Category3Repo.CheckCat3isExist(cat3)
	log.Println("cat3", cat3Existing)
	if cat3Existing == true {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    constants.CAT3_EXISTIG,
		})
	}

	err := svc.Category3Repo.Insert(cat3)

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

func (svc Category3Repo) GetCategory3All(e echo.Context) error {
	listCat3, err := svc.Category3Repo.GetAllCategory3()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, handlers.ErrorResponse{
			HTTPStatus: http.StatusInternalServerError,
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}
	return e.JSON(http.StatusOK, handlers.SuccessResponseList{
		HTTPStatus: http.StatusOK,
		Time:       constants.TIME_NOW,
		Data:       listCat3,
	})
}

func (svc Category3Repo) GetCat3(e echo.Context) error {
	var cat3 = e.Param("id")
	cat3Int, err := strconv.Atoi(cat3)
	cat3Uint := uint(cat3Int)
	if err != nil {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}
	cat3Str, err := svc.Category3Repo.GetCat3ById(cat3Uint)
	if err != nil {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}
	return e.JSON(http.StatusOK, handlers.SuccessReposeDataCat3{
		HTTPStatus: http.StatusOK,
		Time:       constants.TIME_NOW,
		Data:       cat3Str,
	})
}
