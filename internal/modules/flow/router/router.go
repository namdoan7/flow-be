package router

import (
	"go-be/internal/modules/flow/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, h *handler.Handler) {
	orderGroup := r.Group("/flows")
	{
		orderGroup.GET("/", h.GetPaginate)
	}
}
