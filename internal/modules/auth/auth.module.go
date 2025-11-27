package auth

import (
	"go-be/internal/modules/user"

	"github.com/gin-gonic/gin"
)

type Module struct {
	Handler *Handler
	Service Service
}

type Dependencies struct {
	UserRepository user.Repository // interface
}

func NewModule(r *gin.RouterGroup, deps *Dependencies) {
	service := NewService(deps.UserRepository)
	handler := NewHandler(service)

	RegisterRoutes(r, handler)
}
