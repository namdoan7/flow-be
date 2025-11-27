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

func NewModule(deps *Dependencies) *Module {
	repo := NewRepository(deps.DB)
	service := NewService(repo)
	handler := NewHandler(service)

	return &Module{
		Handler:    handler,
		Service:    service,
		Repository: repo,
	}
}

func (m *Module) RegisterRoutes(r *gin.RouterGroup) {
	RegisterRoutes(r, m.Handler)
}
