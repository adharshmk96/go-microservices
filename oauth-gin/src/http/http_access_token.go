package http

import (
	"net/http"
	"strings"

	"github.com/adharshmk96/go-microservices/oauth-gin/src/domain/accesstoken"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
}

type accessTokenHandler struct {
	service accesstoken.Service
}

func NewHandler(service accesstoken.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenID := strings.TrimSpace(c.Param("access_token_id"))

	accessToken, err := h.service.GetById(accessTokenID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}
