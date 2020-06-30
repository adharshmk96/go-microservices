package userdata

import (
	"github.com/adharshmk96/go-microservices/auth-fiber/utils/errors"
)

// Create is used to Create a User
func (user *User) Create() *errors.RestErr {
	// Make an error when db doesnt find the email to check uniqueness
	if Client.Where("email = ?", user.Email).First(&user).RecordNotFound() {
		// Handle by creating the object
		Client.Create(user)
		return nil
	}

	return errors.NewBadRequestError("Email Already Exists")
}

// Get is used to get a user
func (user *User) Get() *errors.RestErr {
	if err := Client.Where("id = ?", user.ID).First(&user).Error; err != nil {
		return errors.NotFoundError("User Not Found !")
	}
	return nil
}