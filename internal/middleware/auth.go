package middleware

import (
	"go-be/internal/common/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			response.Error(c, http.StatusUnauthorized, "Authorization header required", nil)
			c.Abort()
			return
		}

		// TODO: Validate token
		// c.Set("userID", userID)

		c.Next()
	}
}
