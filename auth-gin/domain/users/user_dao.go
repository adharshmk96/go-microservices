package users

// Retrieve and save -> database interaction

import (
	"fmt"

	"github.com/adharshmk96/go-microservices/auth-gin/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

// Get gets user info
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewBadRequestError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

// Save Saves in db
func (user *User) Save() *errors.RestErr {
	if current := usersDB[user.ID]; current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exitsts", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exitsts", user.ID))
	}
	usersDB[user.ID] = user
	return nil
}
