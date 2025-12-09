package router

import (
	"go-be/internal/modules/flow/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, h *handler.Handler) {
	flowGroup := r.Group("/flows")
	{
		flowGroup.GET("/", h.GetPaginate)
		flowGroup.POST("/", h.CreateFlow)
		flowGroup.PUT("/:id", h.UpdateFlow)
	}
}
