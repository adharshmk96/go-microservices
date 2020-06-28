package users

import (
	"strings"

	"github.com/adharshmk96/go-microservices/auth-fiber/utils/errors"
)

// User is the core structure of the data this module handles
type User struct {
	ID          int64  `json:"id" xml:"id" form:"id" query:"id"`
	FirstName   string `json:"first_name" xml:"first_name" form:"first_name" query:"first_name"`
	LastName    string `json:"last_name" xml:"last_name" form:"last_name" query:"last_name"`
	Email       string `json:"email" xml:"email" form:"email" query:"email"`
	DateCreated string `json:"date_created" xml:"date_created" form:"date_created" query:"date_created"`
}

// Validate user inputs
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid Email")
	}
	return nil
}
