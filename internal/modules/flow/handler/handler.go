package handler

import (
	"go-be/internal/common/response"
	"go-be/internal/modules/flow/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetPaginate(c *gin.Context) {
	flows, err := h.service.GetPaginate()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get flows", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Get flows successfully", flows)
}
