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
)

type MaterialsRepo struct {
	MaterialsRepo            repositories.Material
	Category3Repo            repositories.Category3
	MaterialDetailsRepo      repositories.MaterialDetails
	MaterialFieldsDetailRepo repositories.MaterialFieldsDetail
}

func NewMaterials(materialRepo repositories.Material, category3Repo repositories.Category3, materialDetailsRepo repositories.MaterialDetails, materialFieldDetailRepo repositories.MaterialFieldsDetail) *MaterialsRepo {
	return &MaterialsRepo{
		MaterialsRepo:            materialRepo,
		Category3Repo:            category3Repo,
		MaterialDetailsRepo:      materialDetailsRepo,
		MaterialFieldsDetailRepo: materialFieldDetailRepo,
	}
}

func stringPointer(s string) *string {
	return &s
}

func (svc MaterialsRepo) CreateMaterials(e echo.Context) error {
	var req requests.RequestMaterial
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusUnprocessableEntity, handlers.ErrorResponse{
			HTTPStatus: http.StatusUnprocessableEntity,
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}

	if req.Category3Id == 0 {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    "Category3Id is null",
		})
	}
	if req.Code == nil || 0 == len(req.Code) {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    "Code is null",
		})
	}

	cat3Exist, err := svc.Category3Repo.FindCat3ById(req.Category3Id)
	if err != nil {
		return err
	}
	log.Println("Cate3category:", cat3Exist)
	if cat3Exist.Id == 0 {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    "Category3Id does not exist",
		})
	}

	cat3Mapper := dao.Material{
		Category3Id: req.Category3Id,
		Code:        nil,
		Description: nil,
	}

	material, result := svc.MaterialsRepo.Insert(cat3Mapper)
	if result != nil {
		return e.JSON(http.StatusUnprocessableEntity, handlers.ErrorResponse{
			HTTPStatus: http.StatusUnprocessableEntity,
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}
	log.Printf("create success material: %v", material)

	var materialCode string
	var materialDescription string

	for index, code := range req.Code {
		materialDetailMapper := dao.MaterialDetail{
			MaterialId:            material.Id,
			MaterialFieldDetailId: code,
			CodeOrder:             index + 1,
		}

		materialDetail, err := svc.MaterialDetailsRepo.Insert(materialDetailMapper)
		log.Printf("materialDetail: %v", materialDetail)
		if err != nil {
			log.Printf("Error inserting material detail for code %v: %v", code, err)
			continue
		}
		materialFieldDetail := svc.MaterialFieldsDetailRepo.GetMaterialById(materialDetail.MaterialFieldDetailId)
		materialCode += materialFieldDetail.Code
		materialDescription += materialFieldDetail.Name
	}

	updateMaterialMapper := dao.Material{
		Id:          material.Id,
		Code:        stringPointer(cat3Exist.Code + materialCode),
		Description: &materialDescription,
	}

	materials, err := svc.MaterialsRepo.Update(updateMaterialMapper)
	if err != nil {
		log.Printf("Error updating materials: %v", err)
	}

	return e.JSON(http.StatusOK, handlers.SuccessResponseMaterialDetails{
		HTTPStatus: http.StatusOK,
		Time:       constants.TIME_NOW,
		Data:       materials,
	})
}

func (svc MaterialsRepo) GetMaterials(e echo.Context) error {
	material, err := svc.MaterialsRepo.GetALL()
	if err != nil {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}
	return e.JSON(http.StatusOK, handlers.SuccessResponseMaterialList{
		HTTPStatus: http.StatusOK,
		Time:       constants.TIME_NOW,
		Code:       200,
		Data:       material,
	})

}
