package userdata

import (
	"fmt"

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

	panic("Email already Exists")
}

// Get is used to get a user
func (user *User) Get() *errors.RestErr {
	if err := Client.Where("id = ?", user.ID).First(&user).Error; err != nil {
		return errors.NotFoundError("User Not Found !")
	}
	return nil
}

// Update is used to Update a User
func (user *User) Update() *errors.RestErr {
	// Make an error when db doesnt find the email to check uniqueness
	Client.Save(&user)
	return nil
}

// Delete is used to Delete a User
func (user *User) Delete() *errors.RestErr {
	// Make an error when db doesnt find the email to check uniqueness
	Client.Delete(&user)
	return nil
}

// Find User
func (user *User) Find(status string) ([]User, *errors.RestErr) {
	// Select from db
	var users []User
	Client.Where(&user).Find(&users)
	fmt.Println(users)
	return users, nil
}
