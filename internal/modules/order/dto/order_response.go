package dto

import "github.com/google/uuid"

type OrderResponse struct {
	ID     uuid.UUID `json:"id"`
	Total  float64   `json:"total"`
	Status string    `json:"status"`
}
