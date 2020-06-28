package users

// Data Transfer, logic behind being available

import (
	"strings"

	"github.com/adharshmk96/go-microservices/auth-gin/utils/errors"
)

// User is the core structure of the data this module handles
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid Email")
	}
	return nil
}
