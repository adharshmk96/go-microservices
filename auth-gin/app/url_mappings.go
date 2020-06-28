package app

import (
	"github.com/adharshmk96/go-microservices/auth-gin/controllers/ping"
	"github.com/adharshmk96/go-microservices/auth-gin/controllers/users"
)

// mapUrls used to map the url routings
func mapUrls() {
	// Ping Test
	router.GET("/gin/ping", ping.Ping)

	// User Routes
	router.GET("/gin/users/search", users.FindUser)
	router.GET("/gin/user/:user_id", users.GetUser)
	router.POST("/gin/users", users.CreateUser)
}
