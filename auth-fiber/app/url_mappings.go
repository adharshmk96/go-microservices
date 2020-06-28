package app

import (
	"github.com/adharshmk96/go-microservices/auth-fiber/controllers/ping"
	"github.com/adharshmk96/go-microservices/auth-fiber/controllers/users"
)

func mapUrls() {
	// Ping Route
	router.Get("fiber/ping", ping.Ping)

	// User Route
	router.Get("fiber/user/:user_id", users.GetUser)
	router.Post("fiber/users", users.CreateUser)

}
