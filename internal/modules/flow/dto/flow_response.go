package dto

import "github.com/google/uuid"

type FlowResponse struct {
	ID     uuid.UUID `json:"id"`
	Name   float64   `json:"name"`
	Status string    `json:"status"`
}
