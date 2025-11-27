package route

import (
	"go-be/internal/modules/auth"
	"go-be/internal/modules/order"
	"go-be/internal/modules/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Mount(db *gorm.DB) *gin.Engine {
	// User Module
	userModule := user.New(db)

	// Auth Module
	authModule := auth.New(userModule.Repository)

	// Order Module
	orderModule := order.New(db)

	return NewRouter(userModule.Handler, authModule.Handler, orderModule.Handler)
}
