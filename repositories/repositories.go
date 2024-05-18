package repositories

import (
	"daijai-service/models/dao"
	"github.com/google/uuid"
)

type ProjectStatus interface {
	Insert(projectStatus dao.ProjectStatus) error
	GetByProjectId(projectId uuid.UUID) (dao.ProjectStatus, error)
	GetAllProjectStatus() ([]dao.ProjectStatus, error)
}
