package app

import (
	"github.com/adharshmk96/go-microservices/oauth-gin/src/clients/cassandra"
	"github.com/adharshmk96/go-microservices/oauth-gin/src/domain/accesstoken"
	"github.com/adharshmk96/go-microservices/oauth-gin/src/http"
	"github.com/adharshmk96/go-microservices/oauth-gin/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, err := cassandra.GetSession()
	if err != nil {
		panic("Error Creating a session")
	}
	session.Close()

	atService := accesstoken.NewService(db.New())
	atHandler := http.NewHandler(atService)

	router.GET("/gin/oauth/accesstoken/:access_token_id", atHandler.GetById)

	router.Run(":3000")
}
