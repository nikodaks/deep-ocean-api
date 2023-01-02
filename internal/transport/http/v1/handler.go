package v1

import (
	"deep-ocean-api/internal/domain/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {

	return &Handler{services: services}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	h.initAuthRoutes(api)
}
