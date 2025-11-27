package router

import (
	"go-be/internal/modules/user/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, h *handler.Handler) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", h.Register)
	}
}
