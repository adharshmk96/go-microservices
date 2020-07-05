package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/adharshmk96/go-microservices/oauth-gin/src/domain/accesstoken"
	"github.com/adharshmk96/go-microservices/oauth-gin/src/utils/errors"

	"github.com/adharshmk96/go-microservices/oauth-gin/src/service/atservice"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
	// UpdateExpiry(*gin.Context)
}

type accessTokenHandler struct {
	service atservice.Service
}

func NewHandler(service atservice.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenID := strings.TrimSpace(c.Param("access_token_id"))
	fmt.Println("Inside handler")

	accessToken, err := handler.service.GetById(accessTokenID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}
func (handler *accessTokenHandler) Create(c *gin.Context) {
	var at accesstoken.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON")
		c.JSON(restErr.Status, restErr)
		return
	}
	if err := handler.service.Create(at); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, at)

}
