package requests

type RequestProjectStatus struct {
	ProjectName string `json:"projectName"`
	Status      string `json:"status"`
	CreatedBy   string `json:"createdBy"`
}
