package services

import (
	"daijai-service/constants"
	"daijai-service/models/dao"
	"daijai-service/models/handlers"
	"daijai-service/models/requests"
	"daijai-service/models/response"
	"daijai-service/repositories"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ProjectStatusRepo struct {
	ProjectStatusRepo     repositories.ProjectStatus
	EstimateItem          repositories.EstimateItem
	EstimateItemMaterials repositories.EstimateItemMaterial
	Material              repositories.Material
}

func NewProjectStatusService(estimateItemMaterialRepo repositories.EstimateItemMaterial, materialsRepo repositories.Material, estimateItems repositories.EstimateItem, projectRepo repositories.ProjectStatus) *ProjectStatusRepo {
	return &ProjectStatusRepo{
		ProjectStatusRepo:     projectRepo,
		EstimateItem:          estimateItems,
		EstimateItemMaterials: estimateItemMaterialRepo,
		Material:              materialsRepo,
	}
}

func (svc ProjectStatusRepo) CreateProject(e echo.Context) error {
	var r requests.RequestProjectStatus
	if err := e.Bind(&r); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"httpStatus": http.StatusUnprocessableEntity,
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    err.Error(),
		})
	}

	if r.ProjectName == "" {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"httpStatus": http.StatusBadRequest,
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    "ProjectName is empty",
		})
	}

	if r.CreatedBy == "" {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"httpStatus": http.StatusBadRequest,
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    "CreateBy is empty",
		})
	}

	project := dao.Project{
		ProjectName: strings.TrimSpace(strings.ToUpper(r.ProjectName)),
		Status:      "Success Create Project",
		CreatedBy:   strings.TrimSpace(strings.ToUpper(r.CreatedBy)),
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   "",
		Details:     false,
	}

	err := svc.ProjectStatusRepo.Insert(project)

	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"httpStatus": http.StatusUnprocessableEntity,
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    err.Error(),
		})
	}

	// ส่งคืน response พร้อมกับ ID ที่ได้จากการบันทึก
	return e.JSON(http.StatusOK, map[string]interface{}{
		"httpStatus":  http.StatusOK,
		"time":        time.Now().Format("2006-01-02 15:04:05"),
		"projectName": project.ProjectName,
		"status":      "Create Project Suceesss",
	})
}

func (svc *ProjectStatusRepo) GetProjectStatus(e echo.Context) error {
	projectIdStr := e.Param("id")
	projectIdInt, err := strconv.Atoi(projectIdStr)
	if err != nil {
		log.Println(err.Error())
	}
	projectStatus, err := svc.ProjectStatusRepo.GetByProjectId(projectIdInt)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return e.JSON(http.StatusNotFound, map[string]interface{}{
				"httpStatus": http.StatusNotFound,
				"time":       time.Now().Format("2006-01-02 15:04:05"),
				"message":    "Project not found",
			})
		}
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"httpStatus": http.StatusInternalServerError,
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"httpStatus":  http.StatusOK,
		"time":        time.Now().Format("2006-01-02 15:04:05"),
		"projectName": projectStatus.ProjectName,
		"message":     projectStatus,
		"status":      "Fetch Project Success",
	})
}

func (svc ProjectStatusRepo) GetAllProject(e echo.Context) error {
	project, err := svc.ProjectStatusRepo.GetAllProjectStatus()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"httpStatus": http.StatusInternalServerError,
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"httpStatus": http.StatusOK,
		"time":       time.Now().Format("2006-01-02 15:04:05"),
		"project":    project,
	})
}

func (svc ProjectStatusRepo) UpdateProject(e echo.Context) error {
	var rq requests.RequestProjectStatus
	if err := e.Bind(&rq); err != nil {
		return e.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"httpStatus": http.StatusUnprocessableEntity,
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    err,
		})
	}
	if rq.ProjectName == "" {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"httpStatus": http.StatusBadRequest,
			"time":       constants.TIME_NOW,
			"message":    "Project name is empty",
		})
	}
	if rq.CreatedBy == "" {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"httpStatus": http.StatusBadRequest,
			"time":       constants.TIME_NOW,
			"message":    "create by is empty",
		})
	}
	if rq.Status == "" {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"httpStatus": http.StatusBadRequest,
			"time":       constants.TIME_NOW,
			"message":    "Status name is empty",
		})
	}

	project := dao.Project{
		ProjectName: strings.TrimSpace(strings.ToUpper(rq.ProjectName)),
		Status:      strings.TrimSpace(strings.ToUpper(rq.Status)),
		CreatedBy:   strings.TrimSpace(strings.ToUpper(rq.CreatedBy)),
		UpdatedAt:   constants.TIME_NOW,
		Details:     false,
	}

	projectStatus, err := svc.ProjectStatusRepo.UpdateProjectStatus(project)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return e.JSON(http.StatusNotFound, map[string]interface{}{
				"httpStatus": http.StatusNotFound,
				"time":       time.Now().Format("2006-01-02 15:04:05"),
				"message":    "Project not found",
			})
		}
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"httpStatus": http.StatusInternalServerError,
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"httpStatus":  http.StatusOK,
		"time":        time.Now().Format("2006-01-02 15:04:05"),
		"projectName": projectStatus.ProjectName,
		"message":     projectStatus,
		"status":      "update Project Success",
	})

}

func (svc ProjectStatusRepo) DeleteProject(e echo.Context) error {
	projectIdStr := e.Param("projectId")
	projectId, err := uuid.Parse(projectIdStr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"httpStatus": http.StatusBadRequest,
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    "Invalid project ID",
		})
	}
	err = svc.ProjectStatusRepo.DeleteProject(projectId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return e.JSON(http.StatusNotFound, map[string]interface{}{
				"httpStatus": http.StatusNotFound,
				"time":       time.Now().Format("2006-01-02 15:04:05"),
				"message":    "Project not found",
			})
		}
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"httpStatus": http.StatusInternalServerError,
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"http":      http.StatusOK,
		"time":      constants.TIME_NOW,
		"projectId": projectId,
		"message":   "delete success",
	})
}

func (svc *ProjectStatusRepo) GetEstimateItemList(e echo.Context) error {
	projectIdStr := e.Param("ProjectId")
	projectIdInt, err := strconv.Atoi(projectIdStr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}

	projectMaterials := svc.EstimateItemMaterials.FindMaterialById(projectIdInt)
	if len(projectMaterials) == 0 {
		return e.JSON(http.StatusBadRequest, handlers.ErrorResponse{
			HTTPStatus: http.StatusBadRequest,
			Time:       constants.TIME_NOW,
			Message:    "Project not found",
		})
	}

	var responseProject response.ProjectResponseList
	responseProject.Id = projectMaterials[0].Project.Id
	responseProject.ProjectName = projectMaterials[0].Project.ProjectName
	responseProject.CreateBy = projectMaterials[0].Project.CreatedAt

	var itemAll response.EstimateItemAll

	estimateItemType := make(map[int]*response.EstimateItemTypeResponseAll)
	estimateItemsMap := make(map[int]*response.EstimateItemAll)

	for _, material := range projectMaterials {
		estimateItem := &material.EstimateItem

		if _, ok := estimateItemType[estimateItem.EstimateItemType.Id]; !ok {
			estimateItemType[estimateItem.EstimateItemType.Id] = &response.EstimateItemTypeResponseAll{
				Id:   estimateItem.EstimateItemType.Id,
				Name: estimateItem.EstimateItemType.Name,
			}
		}

		if _, ok := estimateItemsMap[estimateItem.Id]; !ok {
			itemAll := response.EstimateItemAll{
				Id:                   estimateItem.Id,
				Name:                 estimateItem.Name,
				Code:                 estimateItem.Code,
				Price:                estimateItem.Price,
				EstimateItemMaterial: []response.EstimateItemMaterial{},
			}
			estimateItemType[estimateItem.EstimateItemType.Id].EstimateItem = append(estimateItemType[estimateItem.EstimateItemType.Id].EstimateItem, itemAll)
		}

		itemMaterial := response.EstimateItemMaterial{
			MaterialAmount: material.MaterialAmount,
			MaterialUnit:   material.MaterialUnit,
			Material:       []dao.Material{material.Material},
		}
		log.Printf("itemMaterial: %s\n", itemMaterial)
		estimateItemsMap[estimateItem.Id].EstimateItemMaterial = append(estimateItemsMap[estimateItem.Id].EstimateItemMaterial, itemMaterial)
		itemAll.EstimateItemMaterial = estimateItemsMap[estimateItem.Id].EstimateItemMaterial
	}

	for _, estimateItemType := range estimateItemType {
		responseProject.Detail = append(responseProject.Detail, *estimateItemType)
	}

	return e.JSON(http.StatusOK, responseProject)
}
