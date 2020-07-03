package main

import (
	"github.com/adharshmk96/go-microservices/auth-fiber/application"
	"github.com/adharshmk96/go-microservices/auth-fiber/models/mysql/userdata"
)

func main() {
	// Default Logger
	// app.Use(middleware.Logger())
	application.StartApplication()

	defer userdata.Client.Close()
}
