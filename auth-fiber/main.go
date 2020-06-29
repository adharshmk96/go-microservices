package main

import (
	"github.com/adharshmk96/go-microservices/auth-fiber/app"
	"github.com/adharshmk96/go-microservices/auth-fiber/datasources/mysql/userdatabase"
)

func main() {
	app.StartApplication()
	defer userdatabase.Client.Close()
}
