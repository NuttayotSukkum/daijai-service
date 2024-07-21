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

	projectMap := make(map[int]*response.ProjectResponseList)
	itemTypeMap := make(map[int]*response.EstimateItemTypeResponseAll)
	estimateItemMap := make(map[int]*response.EstimateItemAll)

	for _, entry := range projectMaterials {
		projectData := entry.Project
		projectId := projectData.Id
		projectName := projectData.ProjectName
		createdBy := projectData.CreatedBy

		itemTypeData := entry.EstimateItem.EstimateItemType
		itemTypeId := itemTypeData.Id
		itemTypeName := itemTypeData.Name

		estimateItemData := entry.EstimateItem
		estimateItemId := estimateItemData.Id
		estimateItemName := estimateItemData.Name
		estimateItemCode := estimateItemData.Code
		estimateItemPrice := estimateItemData.Price
		estimateItemTypeId := estimateItemData.EstimateItemTypeId // Ensure this field exists and is set

		materialData := entry.Material
		materialId := materialData.Id
		cat3 := materialData.Category3
		materialCode := materialData.Code
		materialDescription := materialData.Description
		materialAmount := entry.MaterialAmount
		materialUnit := entry.MaterialUnit

		if _, exists := projectMap[projectId]; !exists {
			projectMap[projectId] = &response.ProjectResponseList{
				Id:          projectId,
				ProjectName: projectName,
				CreateBy:    createdBy,
				Detail:      []response.EstimateItemTypeResponseAll{},
			}
		}

		if _, exists := itemTypeMap[itemTypeId]; !exists {
			itemTypeMap[itemTypeId] = &response.EstimateItemTypeResponseAll{
				Id:           itemTypeId,
				Name:         itemTypeName,
				EstimateItem: []response.EstimateItemAll{},
			}
		}

		if _, exists := estimateItemMap[estimateItemId]; !exists {
			estimateItemMap[estimateItemId] = &response.EstimateItemAll{
				Id:                   estimateItemId,
				Name:                 estimateItemName,
				Code:                 estimateItemCode,
				Price:                estimateItemPrice,
				EstimateItemTypeId:   estimateItemTypeId, // Ensure this field exists and is set
				EstimateItemMaterial: []response.EstimateItemMaterial{},
			}
		}

		estimateItemMaterial := response.EstimateItemMaterial{
			MaterialAmount: materialAmount,
			MaterialUnit:   materialUnit,
			Material: []dao.Material{{
				Id:          materialId,
				Category3:   cat3,
				Code:        materialCode,
				Description: materialDescription,
			}},
		}

		estimateItemMap[estimateItemId].EstimateItemMaterial = append(estimateItemMap[estimateItemId].EstimateItemMaterial, estimateItemMaterial)
	}

	for _, item := range estimateItemMap {
		itemTypeId := item.EstimateItemTypeId
		if _, exists := itemTypeMap[itemTypeId]; exists {
			itemTypeMap[itemTypeId].EstimateItem = append(itemTypeMap[itemTypeId].EstimateItem, *item)
		} else {
			log.Println("Item type not found:", itemTypeId)
		}
		log.Println("Step 12")
	}
	log.Println("Step 20")
	for _, itemType := range itemTypeMap {
		for projectId := range projectMap {
			if itemType == nil {
				log.Printf("projectId:%d", projectId)
			} else {
				projectMap[projectId].Detail = append(projectMap[projectId].Detail, *itemType)
				log.Printf("items:%s", projectMap[projectId].Detail)
			}
		}
	}

	var projectResponseList []response.ProjectResponseList
	for _, project := range projectMap {
		projectResponseList = append(projectResponseList, *project)
	}

	return e.JSON(http.StatusOK, projectResponseList)
}
