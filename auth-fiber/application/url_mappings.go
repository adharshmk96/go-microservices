package application

import (
	"github.com/adharshmk96/go-microservices/auth-fiber/controllers/ping"
	"github.com/adharshmk96/go-microservices/auth-fiber/controllers/users"
)

func mapUrls() {
	// Ping Route
	app.Get("fiber/ping", ping.Ping)

	// User Route
	app.Post("fiber/users", users.CreateUser)
	app.Get("fiber/user/:user_id", users.GetUser)
	app.Put("fiber/user/:user_id", users.UpdateUser)
	app.Patch("fiber/user/:user_id", users.UpdateUser)
	app.Delete("fiber/user/:user_id", users.DeleteUser)

}
