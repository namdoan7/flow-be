package dto

type CreateOrderRequest struct {
	Total float64 `json:"total" binding:"required,gt=0"`
}
