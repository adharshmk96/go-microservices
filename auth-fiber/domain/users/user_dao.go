package users

import (
	"fmt"

	"github.com/adharshmk96/go-microservices/auth-fiber/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

// Get user info
func (user *User) Get() *errors.RestErr {
	result := userDB[user.ID]
	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

// Save user info
func (user *User) Save() *errors.RestErr {
	if current := userDB[user.ID]; current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exitsts", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exitsts", user.ID))
	}
	userDB[user.ID] = user
	return nil
}
