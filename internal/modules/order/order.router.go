package order

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	orderGroup := r.Group("/orders")
	{
		orderGroup.POST("/", h.CreateOrder)
	}
}
