package repositories

import (
	"daijai-service/models/dao"
	"github.com/google/uuid"
)

type ProjectStatus interface {
	Insert(projectStatus dao.ProjectStatus) error
	GetByProjectId(projectId uuid.UUID) (dao.ProjectStatus, error)
	GetAllProjectStatus() ([]dao.ProjectStatus, error)
	UpdateProjectStatus(project dao.ProjectStatus) (dao.ProjectStatus, error)
	DeleteProject(projectId uuid.UUID) error
}

type Category3 interface {
	Insert(category3 dao.Category3) error
	CheckCat3isExist(req dao.Category3) bool
	GetAllCategory3() ([]dao.Category3, error)
	GetCat3ById(id uint) (dao.Category3, error)
	//GetCat3ById(id uuid.UUID) dao.Category3
}

type MaterialFields interface {
	Insert(req dao.MaterialFields) error
	CheckExisting(req dao.MaterialFields) bool
	GetMaterialAll() []dao.MaterialFields
	GetMaterialById(id uint) dao.MaterialFields
}

type MaterialFieldsDetail interface {
	Insert(req dao.MaterialFieldDetails) error
	CheckExisting(Name string) bool
	GetMaterialAll() []dao.MaterialFieldDetails
	GetMaterialById(id uint) dao.MaterialFieldDetails
	GetMaterialByMaterialFields(id string) dao.MaterialFieldDetails
}
