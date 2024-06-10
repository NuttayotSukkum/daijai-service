package services

import "github.com/labstack/echo/v4"

type ProjectStatus interface {
	CreateProject(e echo.Context) error
	GetProjectStatus(e echo.Context) error
	GetAllProject(e echo.Context) error
	UpdateProject(e echo.Context) error
	DeleteProject(e echo.Context) error
}

type Category3 interface {
	CreateCategory3(e echo.Context) error
	GetCategory3All(e echo.Context) error
	GetCat3(e echo.Context) error
}

type Material interface {
	CreateMaterials(e echo.Context) error
}
