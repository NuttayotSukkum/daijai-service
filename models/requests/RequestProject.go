package requests

type RequestProjectStatus struct {
	ProjectName string `json:"project_name"`
	CreatedBy   string `json:"created_by"`
	Details     string `json:"details"`
}
