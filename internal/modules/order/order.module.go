package order

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	Handler    *Handler
	Service    Service
	Repository Repository
}

type Dependencies struct {
	DB *gorm.DB
}

func NewModule(r *gin.RouterGroup, deps *Dependencies) {
	repo := NewRepository(deps.DB)
	service := NewService(repo)
	handler := NewHandler(service)

	RegisterRoutes(r, handler)
}
