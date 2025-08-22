package v1

import (
	"deep-ocean-api/pkg/logger"

	"github.com/gin-gonic/gin"
)

type response struct {
	Message string `json:"message"`
}

func errorResponse(c *gin.Context, statusCode int, message string) {
	logger.Error(message)
	c.AbortWithStatusJSON(statusCode, response{message})
}
