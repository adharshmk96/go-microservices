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
	router.GET("/gin/user/:user_id", users.Get)
	router.GET("/gin/internal/users/search", users.Search)
	router.POST("/gin/users", users.Create)
	router.PUT("/gin/user/:user_id", users.Update)
	router.PATCH("/gin/user/:user_id", users.Update)
	router.DELETE("/gin/user/:user_id", users.Delete)
}
