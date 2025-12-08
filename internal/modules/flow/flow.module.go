package flow

import (
	"go-be/internal/modules/flow/handler"
	"go-be/internal/modules/flow/repository"
	"go-be/internal/modules/flow/router"
	"go-be/internal/modules/flow/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	Handler    *handler.Handler
	Service    service.Service
	Repository repository.Repository
}

type Dependencies struct {
	DB *gorm.DB
}

func NewModule(r *gin.RouterGroup, deps *Dependencies) {
	repo := repository.NewRepository(deps.DB)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	router.RegisterRoutes(r, h)
}
