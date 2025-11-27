package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", h.Register)
	}
}
