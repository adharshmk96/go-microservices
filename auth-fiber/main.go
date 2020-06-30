package main

import (
	"github.com/adharshmk96/go-microservices/auth-fiber/app"
	"github.com/adharshmk96/go-microservices/auth-fiber/models/mysql/userdata"
)

func main() {
	app.StartApplication()
	defer userdata.Client.Close()
}
