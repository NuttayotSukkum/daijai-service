package db

import (
	"daijai-service/models/dao"
	"daijai-service/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (repo ProjectStatus) GetByProjectId(projectId uuid.UUID) (dao.ProjectStatus, error) {
	var projectStatus dao.ProjectStatus
	err := repo.db.Where("project_id = ?", projectId).First(&projectStatus).Error
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
