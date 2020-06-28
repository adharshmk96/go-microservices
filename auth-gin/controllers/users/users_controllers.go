package users

import (
	"net/http"

	"github.com/adharshmk96/go-microservices/auth-gin/domain/users"
	"github.com/adharshmk96/go-microservices/auth-gin/services"
	"github.com/adharshmk96/go-microservices/auth-gin/utils/errors"
	"github.com/gin-gonic/gin"
)

// CreateUser creates a user / registers
func CreateUser(c *gin.Context) {
	// Make a struct
	var user users.User

	// Behind the scnenes
	// // get data from request body
	// bytes, err := ioutil.ReadAll(c.Request.Body)

	// if err != nil {
	// 	//TODO: Handle Error
	// 	return
	// }

	// // Parse Json with user type
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	//TODO: Handle Json Error
	// 	return
	// }
	if err := c.ShouldBindJSON(&user); err != nil {
		// Handle JSON Error
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	result, validErr := services.CreateUser(user)
	if validErr != nil {
		//TODO: Handle user creation err
		c.JSON(validErr.Status, validErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// GetUser returns the user info
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Me")
}

// FindUser is sued to search the user Info
func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Me")
}
