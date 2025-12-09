package handler

import (
	"go-be/internal/common/response"
	"go-be/internal/modules/flow/dto"
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

func (h *Handler) CreateFlow(c *gin.Context) {
	var req dto.CreateFlowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

}

func (h *Handler) UpdateFlow(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateFlowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}
	flow, err := h.service.UpdateFlow(id, req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "error", err.Error())
		return
	}
	if flow == nil {
		response.Error(c, http.StatusNotFound, "Not found", gin.H{"message": "flow not found"})
		return
	}

	response.Success(c, http.StatusOK, "Update flow successfully", flow)
}
