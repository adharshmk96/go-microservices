package main

import (
	"github.com/adharshmk96/go-microservices/auth-fiber/application"
	"github.com/adharshmk96/go-microservices/auth-fiber/models/mysql/userdata"
)

func main() {
	application.StartApplication()

	defer userdata.Client.Close()
}
