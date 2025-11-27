package router

import (
	"go-be/internal/modules/order/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, h *handler.Handler) {
	orderGroup := r.Group("/orders")
	{
		orderGroup.POST("/", h.CreateOrder)
	}
}
