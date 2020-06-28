package users

import (
	"net/http"
	"strconv"

	"github.com/adharshmk96/go-microservices/auth-fiber/services"
	"github.com/adharshmk96/go-microservices/auth-fiber/utils/errors"

	"github.com/adharshmk96/go-microservices/auth-fiber/domain/users"
	"github.com/gofiber/fiber"
)

// GetUser Returns current user
func GetUser(c *fiber.Ctx) {
	userID, userErr := strconv.ParseInt(c.Params("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid User id")
		c.Status(err.Status)
		c.JSON(err)
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.Send(getErr.Status)
		c.JSON(getErr)
		return
	}

	c.Status(http.StatusOK)
	c.JSON(user)
}

// CreateUser used for Registration
func CreateUser(c *fiber.Ctx) {

	// Define the Structure
	user := new(users.User)

	if err := c.BodyParser(user); err != nil {
		// TODO: Handle Error
		jsonError := errors.NewBadRequestError("Json Format Error !")
		c.Status(jsonError.Status)
		c.JSON(jsonError)
		return
	}

	result, validErr := services.CreateUser(*user)

	if validErr != nil {
		// TODO: Handle user createion error
		c.Status(validErr.Status)
		c.JSON(validErr)
		return
	}

	c.Status(http.StatusCreated)
	c.JSON(result)
}
