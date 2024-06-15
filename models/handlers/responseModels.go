package handlers

import (
	"daijai-service/models/dao"
	"daijai-service/models/response"
)

type ErrorResponse struct {
	HTTPStatus int    `json:"httpStatus"`
	Time       string `json:"time"`
	Message    string `json:"message"`
}

type SuccessResponse struct {
	HTTPStatus  int         `json:"httpStatus"`
	Time        string      `json:"time"`
	ProjectName string      `json:"projectName"`
	Message     interface{} `json:"message"`
	Status      string      `json:"status"`
}

type SuccessReposeMessage struct {
	HTTPStatus int    `json:"httpStatus"`
	Time       string `json:"time"`
	Name       string `json:"name"`
	Message    string `json:"message"`
}

type SuccessResponseList struct {
	HTTPStatus int             `json:"httpStatus"`
	Time       string          `json:"time"`
	Data       []dao.Category3 `json:"data"`
}

type SuccessReposeDataCat3 struct {
	HTTPStatus int           `json:"httpStatus"`
	Time       string        `json:"time"`
	Data       dao.Category3 `json:"Data"`
}

type SuccessResponseListMaterialFields struct {
	HTTPStatus int                 `json:"httpStatus"`
	Time       string              `json:"time"`
	Data       []dao.MaterialField `json:"data"`
}

type SuccessResponseMaterialFields struct {
	HTTPStatus int               `json:"httpStatus"`
	Time       string            `json:"time"`
	Data       dao.MaterialField `json:"data"`
}

type SuccessResponseMaterialFieldDetailsList struct {
	HTTPStatus int                       `json:"httpStatus"`
	Time       string                    `json:"time"`
	Data       []dao.MaterialFieldDetail `json:"data"`
}

type SuccessResponseMaterialFieldDetails struct {
	HTTPStatus int                     `json:"httpStatus"`
	Time       string                  `json:"time"`
	Data       dao.MaterialFieldDetail `json:"data"`
}

type SuccessResponseMaterialDetails struct {
	HTTPStatus int          `json:"httpStatus"`
	Time       string       `json:"time"`
	Data       dao.Material `json:"data"`
}

type SuccessResponseMaterialList struct {
	HTTPStatus int            `json:"httpStatus"`
	Time       string         `json:"time"`
	Code       int            `json:"code"`
	Data       []dao.Material `json:"data"`
}

type SuccessResponseEstimateItemTypeList struct {
	HTTPStatus int                     `json:"httpStatus"`
	Time       string                  `json:"time"`
	Data       []dao.EstimateItemTypes `json:"data"`
}

type SuccessResponseEstimateItemList struct {
	HTTPStatus int                `json:"httpStatus"`
	Time       string             `json:"time"`
	Data       []dao.EstimateItem `json:"data"`
}

type SuccessResponseEstimateItemMaterial struct {
	HTTPStatus int                               `json:"httpStatus"`
	Time       string                            `json:"time"`
	Data       response.EstimateItemTypeResponse `json:"data"`
}
