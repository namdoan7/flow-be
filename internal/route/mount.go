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

	// User Module
	userModule := user.NewModule(&user.Dependencies{DB: db})

	// Auth Module
	authModule := auth.NewModule(userModule.Repository)

	// Order Module
	orderModule := order.NewModule(&order.Dependencies{DB: db})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	api := r.Group("/api")
	{
		// Protected routes
		// protected := api.Group("/")
		// protected.Use(middleware.AuthMiddleware())

		authModule.RegisterRoutes(api)
		userModule.RegisterRoutes(api)
		orderModule.RegisterRoutes(api)
	}

	return r
}
