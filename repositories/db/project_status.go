package db

import (
	"daijai-service/models/dao"
	"daijai-service/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type ProjectStatus struct {
	db *gorm.DB
}

func NewProjectStatusRepository(db *gorm.DB) repositories.ProjectStatus {
	return &ProjectStatus{db: db}
}

func (repo ProjectStatus) Insert(projectStatus dao.ProjectStatus) error {
	err := repo.db.Create(&projectStatus).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo ProjectStatus) GetByProjectName(projectName string) (dao.ProjectStatus, error) {
	var projectStatus dao.ProjectStatus
	err := repo.db.Where("project_name = ?", projectName).First(&projectStatus).Error
	if err != nil {
		return projectStatus, err
	}
	return projectStatus, nil
}

func (repo ProjectStatus) GetAllProjectStatus() ([]dao.ProjectStatus, error) {
	var allProject []dao.ProjectStatus
	err := repo.db.Find(&allProject).Error
	if err != nil {
		return nil, err
	}
	return allProject, nil
}

func (repo *ProjectStatus) UpdateProjectStatus(project dao.ProjectStatus) (dao.ProjectStatus, error) {
	var response dao.ProjectStatus

	if err := repo.db.Where("project_name = ?", project.ProjectName).First(&response).Error; err != nil {
		return response, err
	}
	response.ProjectName = project.ProjectName
	response.Status = project.Status
	response.UpdatedAt = project.UpdatedAt
	response.CreatedBy = project.CreatedBy
	response.Details = project.Details

	if err := repo.db.Save(&response).Error; err != nil {
		log.Printf("Error updating project status: %v", err)
		return response, err
	}

	return response, nil
}

func (repo *ProjectStatus) DeleteProject(projectId uuid.UUID) error {
	var project dao.ProjectStatus
	err := repo.db.Where("project_id = ?", projectId).Delete(project).Error
	if err != nil {
		return err
	}
	return nil
}
