package route

import (
	"go-be/internal/middleware"
	"go-be/internal/modules/auth"
	"go-be/internal/modules/order"
	"go-be/internal/modules/user"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	userHandler *user.Handler,
	authHandler *auth.Handler,
	orderHandler *order.Handler,
) *gin.Engine {
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	// Public routes
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	api := r.Group("/api")
	{
		user.RegisterRoutes(api, userHandler)
		auth.RegisterRoutes(api, authHandler)

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		order.RegisterRoutes(protected, orderHandler)
	}

	return r
}
