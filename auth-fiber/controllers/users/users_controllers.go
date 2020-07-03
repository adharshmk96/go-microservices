package users

import (
	"net/http"
	"strconv"

	"github.com/adharshmk96/go-microservices/auth-fiber/models/mysql/userdata"
	"github.com/adharshmk96/go-microservices/auth-fiber/services"
	"github.com/adharshmk96/go-microservices/auth-fiber/utils/errors"

	"github.com/gofiber/fiber"
)

// getUserID extracts andr eturn user id
func getUserID(userID string) (uint64, *errors.RestErr) {
	uid, userErr := strconv.ParseUint(userID, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("invalid User id")
	}
	return uid, nil
}

// GetUser Returns current user
func GetUser(c *fiber.Ctx) {
	userID, userErr := getUserID(c.Params("user_id"))
	if userErr != nil {
		c.Status(userErr.Status)
		panic("Invalid User ID")
		// c.JSON(userErr)
		// return
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
	user := new(userdata.User)

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

// UpdateUser updates a user record
func UpdateUser(c *fiber.Ctx) {
	userID, userErr := getUserID(c.Params("user_id"))
	if userErr != nil {
		c.Status(userErr.Status)
		c.JSON(userErr)
		return
	}

	user := new(userdata.User)

	if err := c.BodyParser(user); err != nil {
		jsonError := errors.NewBadRequestError("Json Format Error")
		c.Status(jsonError.Status)
		c.JSON(jsonError)
		return
	}

	user.ID = userID

	isPartial := c.Method() == http.MethodPatch

	result, validErr := services.UpdateUser(isPartial, *user)
	if validErr != nil {
		// TODO: Handle user createion error
		c.Status(validErr.Status)
		c.JSON(validErr)
		return
	}
	// user
	c.Status(http.StatusCreated)
	c.JSON(result)
}

// DeleteUser deletes a user record
func DeleteUser(c *fiber.Ctx) {
	userID, userErr := getUserID(c.Params("user_id"))
	if userErr != nil {
		c.Status(userErr.Status)
		c.JSON(userErr)
		return
	}

	if err := services.DeleteUser(userID); err != nil {
		c.Status(userErr.Status)
		c.JSON(userErr)
		return
	}

	c.Status(http.StatusOK)
	c.JSON(map[string]string{"Status": "success"})
}
