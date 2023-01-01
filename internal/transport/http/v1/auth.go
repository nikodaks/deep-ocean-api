package v1

import (
	"deep-ocean-api/internal/domain/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initAuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
	}
}

func (h *Handler) signUp(c *gin.Context) {
	var dto dto.SignUpDto

	if err := c.BindJSON(&dto); err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	if err := h.services.Auth.SignUp(&dto); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
	}
}
