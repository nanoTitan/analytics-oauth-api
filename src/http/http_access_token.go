package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nanoTitan/analytics-oauth-api/src/domain/atdomain"
	"github.com/nanoTitan/analytics-oauth-api/src/services/accesstoken"
	"github.com/nanoTitan/analytics-oauth-api/src/utils/errors"
)

// AccessTokenHandler - interface for a generic access token handler
type AccessTokenHandler interface {
	GetByID(*gin.Context)
	Create(c *gin.Context)
}

type accessTokenHandler struct {
	service accesstoken.Service
}

// NewAccessTokenHandler - get a new accessTokenHandler object
func NewAccessTokenHandler(service accesstoken.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

// GetByID - create an http request to get an existing access token by user ID
func (handler *accessTokenHandler) GetByID(c *gin.Context) {
	accessTokenID := strings.TrimSpace(c.Param("access_token_id"))

	accessToken, err := handler.service.GetByID(accessTokenID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

// Create - Creates an http request for new access token
func (handler *accessTokenHandler) Create(c *gin.Context) {
	var request atdomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	accessToken, err := handler.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
