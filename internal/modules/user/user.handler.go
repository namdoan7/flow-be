package user

import (
	"go-be/internal/common/response"
	"go-be/internal/modules/user/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	if err := h.service.Register(req.Name, req.Email, req.Password); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to register user", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "User registered successfully", nil)
}
