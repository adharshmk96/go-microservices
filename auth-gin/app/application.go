package app

import (
	"github.com/adharshmk96/go-microservices/auth-gin/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication is the pre entry point function
func StartApplication() {
	mapUrls()
	logger.Info("About to start Application")
	router.Run(":3001")
}
