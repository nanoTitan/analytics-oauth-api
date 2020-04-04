package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nanoTitan/analytics-oauth-api/src/clients/cassandra"
	"github.com/nanoTitan/analytics-oauth-api/src/domain/accesstoken"
	"github.com/nanoTitan/analytics-oauth-api/src/http"
	"github.com/nanoTitan/analytics-oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

// StartApplication - start of the OAuth application
func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()

	atHandler := http.NewHandler(accesstoken.NewService(db.New()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
