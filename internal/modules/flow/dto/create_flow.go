package dto

type CreateFlowRequest struct {
	Total float64 `json:"total" binding:"required,gt=0"`
}
