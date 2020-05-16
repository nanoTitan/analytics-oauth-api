package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nanoTitan/analytics-oauth-api/src/http"
	"github.com/nanoTitan/analytics-oauth-api/src/repository/db"
	"github.com/nanoTitan/analytics-oauth-api/src/repository/rest"
	"github.com/nanoTitan/analytics-oauth-api/src/services/accesstoken"
)

var (
	router = gin.Default()
)

// StartApplication - start of the OAuth application
func StartApplication() {
	atHandler := http.NewAccessTokenHandler(
		accesstoken.NewService(rest.NewRepository(), db.New()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
