package repositories

import (
	"daijai-service/models/dao"
	"github.com/google/uuid"
)

type ProjectStatus interface {
	Insert(projectStatus dao.Project) error
	GetByProjectId(projectId int) (dao.Project, error)
	GetAllProjectStatus() ([]dao.Project, error)
	UpdateProjectStatus(project dao.Project) (dao.Project, error)
	DeleteProject(projectId uuid.UUID) error
}

type Category3 interface {
	Insert(category3 dao.Category3) error
	CheckCat3isExist(req dao.Category3) bool
	GetAllCategory3() ([]dao.Category3, error)
	GetCat3ById(id int) (dao.Category3, error)
	FindCat3ById(id int) (dao.Category3, error)
	//GetCat3ById(id uuid.UUID) dao.Category3
}

type MaterialFields interface {
	Insert(req dao.MaterialField) error
	CheckExisting(req dao.MaterialField) bool
	GetMaterialAll() []dao.MaterialField
	GetMaterialById(id int) dao.MaterialField
}

type MaterialFieldsDetail interface {
	Insert(req dao.MaterialFieldDetail) error
	CheckExisting(Name string) bool
	GetMaterialAll() []dao.MaterialFieldDetail
	GetMaterialById(id int) dao.MaterialFieldDetail
	GetMaterialByMaterialFields(id string) dao.MaterialFieldDetail
}

type Material interface {
	//Insert(material dao.Material) (dao.Material, error)
	Insert(material dao.Material) (dao.Material, error)
	Update(material dao.Material) (dao.Material, error)
	GetALL() ([]dao.Material, error)
	FindMaterialsExist(id int) []dao.Material
	FindMaterialsById(id int) dao.Material
}

type MaterialDetails interface {
	Insert(material dao.MaterialDetail) (dao.MaterialDetail, error)
}

type EstimateItemType interface {
	Insert(estimateItemType dao.EstimateItemTypes) error
	FindEstimateItemTypesByName(name string) dao.EstimateItemTypes
	FindEstimateItemTypesById(id int) ([]dao.EstimateItemTypes, error)
	FindEstimateItemTypeAll() ([]dao.EstimateItemTypes, error)
	FindEstimateItemsTypeById(id int) (dao.EstimateItemTypes, error)
}

type EstimateItem interface {
	Insert(estimateItem dao.EstimateItem) error
	FindEstimateItemAll() ([]dao.EstimateItem, error)
	FindEstimateItemExist(id int) []dao.EstimateItem
	FindEstimateItemById(id int) (dao.EstimateItem, error)
}

type EstimateItemMaterial interface {
	Insert(estimateItemMaterial dao.EstimateItemMaterial) (dao.EstimateItemMaterial, error)
	FindMaterialById(projectId int) []dao.EstimateItemMaterial
}

type ProjectEstimateItem interface {
	Insert(projectEstimateItem dao.ProjectEstimateItem) error
}
