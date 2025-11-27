package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Dependencies struct {
	DB *gorm.DB
}

func NewModule(r *gin.RouterGroup, deps *Dependencies) {
	repo := NewRepository(deps.DB)
	service := NewService(repo)
	handler := NewHandler(service)

	RegisterRoutes(r, handler)
}
