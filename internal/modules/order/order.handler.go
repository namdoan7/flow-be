package order

import (
	"go-be/internal/common/response"
	"go-be/internal/modules/order/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var req dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	// TODO: Get UserID from context (middleware)
	userID := uuid.New() // Mock

	if err := h.service.CreateOrder(userID, req.Total); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create order", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Order created successfully", nil)
}
