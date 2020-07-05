package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication is the pre entry point function
func StartApplication() {
	mapUrls()
	router.Run(":3001")
}
