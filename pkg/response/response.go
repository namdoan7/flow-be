package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-be/internal/repository"
)

type Envelope struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func Success(c *gin.Context, status int, data any) {
	if status == http.StatusNoContent {
		c.Status(status)
		return
	}
	c.JSON(status, Envelope{Data: data})
}

func Error(c *gin.Context, status int, err error) {
	c.JSON(status, Envelope{Error: err.Error()})
}

func HandleDomainError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, repository.ErrUserNotFound):
		Error(c, http.StatusNotFound, err)
	default:
		Error(c, http.StatusInternalServerError, err)
	}
}
