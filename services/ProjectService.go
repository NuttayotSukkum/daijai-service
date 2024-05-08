package services

import (
	"daijai-service/models/dao"
	"daijai-service/models/requests"
	"daijai-service/orm"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
)

func CeateProject(e echo.Context) error {
	var r requests.RequestProject
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

	if r.CreateBy == "" {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"httpStatus": http.StatusBadRequest,
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    "CreateBy is empty",
		})
	}

	project := dao.Project{
		ProjectName: strings.TrimSpace(strings.ToUpper(r.ProjectName)),
		CreatedBy:   strings.TrimSpace(strings.ToUpper(r.CreateBy)),
	}

	if err := orm.Save(&project).Error; err != nil {
		return e.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"httpStatus": http.StatusUnprocessableEntity,
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
