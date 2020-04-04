package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nanoTitan/analytics-oauth-api/src/domain/accesstoken"
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

// NewHandler - get a new accessTokenHandler object
func NewHandler(service accesstoken.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetByID(c *gin.Context) {
	accessTokenID := strings.TrimSpace(c.Param("access_token_id"))

	accessToken, err := h.service.GetByID(accessTokenID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var at accesstoken.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := h.service.Create(at); err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusCreated, at)
}
