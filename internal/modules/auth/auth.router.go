package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", h.Login)
	}
}
