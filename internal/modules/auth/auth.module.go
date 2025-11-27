package auth

import (
	"go-be/internal/modules/user"

	"github.com/gin-gonic/gin"
)

type Module struct {
	Handler *Handler
	Service Service
}

func NewModule(userRepo user.Repository) *Module {
	service := NewService(userRepo)
	handler := NewHandler(service)

	return &Module{
		Handler: handler,
		Service: service,
	}
}

func (m *Module) RegisterRoutes(r *gin.RouterGroup) {
	RegisterRoutes(r, m.Handler)
}
