package services

import (
	"daijai-service/models/dao"
	"daijai-service/models/requests"
	"daijai-service/repositories"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

type ProjectStatusRepo struct {
	ProjectStatusRepo repositories.ProjectStatus
}

func NewProjectStatusService(repo repositories.ProjectStatus) *ProjectStatusRepo {
	return &ProjectStatusRepo{ProjectStatusRepo: repo}
}

func (svc ProjectStatusRepo) CreateProject(e echo.Context) error {
	var r requests.RequestProjectStatus
	if err := e.Bind(&r); err != nil {
		return e.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
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

	project := dao.ProjectStatus{
		ProjectId:   uuid.New(),
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
	projectIdStr := e.Param("projectId")
	projectId, err := uuid.Parse(projectIdStr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"httpStatus": http.StatusBadRequest,
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    "Invalid project ID",
		})
	}

	projectStatus, err := svc.ProjectStatusRepo.GetByProjectId(projectId)
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
