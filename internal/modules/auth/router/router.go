package router

import (
	"go-be/internal/modules/auth/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, h *handler.Handler) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", h.Login)
	}
}
