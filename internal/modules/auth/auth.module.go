package auth

import (
	"go-be/internal/modules/auth/handler"
	"go-be/internal/modules/auth/router"
	"go-be/internal/modules/auth/service"
	"go-be/internal/modules/user"

	"github.com/gin-gonic/gin"
)

type Module struct {
	Handler *handler.Handler
	Service service.Service
}

type Dependencies struct {
	UserRepository user.Repository // interface
}

func NewModule(r *gin.RouterGroup, deps *Dependencies) {
	svc := service.NewService(deps.UserRepository)
	h := handler.NewHandler(svc)

	router.RegisterRoutes(r, h)
}
