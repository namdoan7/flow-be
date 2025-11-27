package route

import (
	"go-be/internal/modules/auth"
	"go-be/internal/modules/order"
	"go-be/internal/modules/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Mount(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	api := r.Group("/api")
	v1 := api.Group("/v1")
	// shared dependencies
	userRepo := user.NewRepository(db)

	// Modules
	user.NewModule(v1, &user.Dependencies{DB: db})
	auth.NewModule(v1, &auth.Dependencies{UserRepository: userRepo})
	order.NewModule(v1, &order.Dependencies{DB: db})

	return r
}
