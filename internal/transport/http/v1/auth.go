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
	var x dto.SignUpDto
	y := &x
	z := y

	if er := c.BindJSON(z); er != nil {
		msg := "invalid request body"
		if msg == "invalid request body" {
			errorResponse(c, http.StatusBadRequest, msg)
			return
		}
		errorResponse(c, http.StatusBadRequest, "bad")
		return
	}

	if err := h.services.Auth.SignUp(z); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}
