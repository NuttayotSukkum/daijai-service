package repositories

import (
	"daijai-service/models/dao"
	"github.com/google/uuid"
)

type ProjectStatus interface {
	Insert(projectStatus dao.ProjectStatus) error
	GetByProjectName(projectName string) (dao.ProjectStatus, error)
	GetAllProjectStatus() ([]dao.ProjectStatus, error)
	UpdateProjectStatus(project dao.ProjectStatus) (dao.ProjectStatus, error)
	DeleteProject(projectId uuid.UUID) error
}
