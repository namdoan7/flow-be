package dto

type UpdateFlowRequest struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	DataString string `json:"data_string"`
}
