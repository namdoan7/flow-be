package user

import (
	"go-be/internal/modules/user/handler"
	"go-be/internal/modules/user/model"
	"go-be/internal/modules/user/repository"
	"go-be/internal/modules/user/router"
	"go-be/internal/modules/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Export types and interfaces for use by other modules
type User = model.User
type Repository = repository.Repository

func NewRepository(db *gorm.DB) Repository {
	return repository.NewRepository(db)
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
