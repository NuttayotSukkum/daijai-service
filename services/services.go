package services

import "github.com/labstack/echo/v4"

type ProjectStatus interface {
	CreateProject(e echo.Context) error
	GetProjectStatus(e echo.Context) error
}
