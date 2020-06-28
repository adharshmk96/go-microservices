package users

import (
	"fmt"
	"net/http"

	"github.com/adharshmk96/go-microservices/auth-fiber/services"

	"github.com/adharshmk96/go-microservices/auth-fiber/domain/users"
	"github.com/gofiber/fiber"
)

// GetUser Returns current user
func GetUser(c *fiber.Ctx) {
	c.Status(http.StatusNotImplemented)
	c.SendString("Implement Me")
}

// CreateUser used for Registration
func CreateUser(c *fiber.Ctx) {

	// Define the Structure
	user := new(users.User)

	if err := c.BodyParser(user); err != nil {
		// TODO: Handle Error
		fmt.Println("Error", err.Error())
		return
	}

	result, validErr := services.CreateUser(*user)

	if validErr != nil {
		// TODO: Handle user createion error
		fmt.Println("Error", validErr.Error())
		return
	}

	c.Status(http.StatusCreated)
	c.JSON(result)
}

// FindUser used for Finding user by Id
func FindUser(c *fiber.Ctx) {
	c.Status(http.StatusNotImplemented)
	c.SendString("Implement Me")
}
