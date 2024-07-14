package requests

type RequestProjectEstimateItem struct {
	ProjectId      int   `json:"projectId"`
	EstimateItemId []int `json:"estimateItemId"`
}
