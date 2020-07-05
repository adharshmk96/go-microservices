package users

import (
	"net/http"
	"strconv"

	"github.com/adharshmk96/go-microservices/auth-gin/domain/users"
	"github.com/adharshmk96/go-microservices/auth-gin/services"
	"github.com/adharshmk96/go-microservices/auth-gin/utils/errors"
	"github.com/gin-gonic/gin"
)

// getUserId used to get the user
func getUserID(userIDParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIDParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("Invalid user id")
	}
	return userID, nil
}

// Create creates a user / registers
func Create(c *gin.Context) {
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

	result, validErr := services.UsersService.CreateUser(user)
	if validErr != nil {
		//TODO: Handle user creation err
		c.JSON(validErr.Status, validErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// Get returns the user info
func Get(c *gin.Context) {
	userID, userErr := getUserID(c.Param("user_id"))

	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	user, getErr := services.UsersService.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

// Update updates the user given the id
func Update(c *gin.Context) {
	userID, userErr := getUserID(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	// userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	result, validErr := services.UsersService.UpdateUser(isPartial, user)
	if validErr != nil {
		c.JSON(validErr.Status, validErr)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))

}

// Delete used to delete user
func Delete(c *gin.Context) {
	userID, userErr := getUserID(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	if err := services.UsersService.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"Status": "deleted"})
}

// Search returns user
func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UsersService.FindUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	result := users.Marshall(c.GetHeader("X-Public") == "true")

	c.JSON(http.StatusOK, result)
}
