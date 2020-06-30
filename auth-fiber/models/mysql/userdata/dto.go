package userdata

import (
	"strings"
	"time"

	"github.com/adharshmk96/go-microservices/auth-fiber/utils/errors"
)

// User describes a User
type User struct {
	ID        uint64 `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Email     string `gorm:"type:varchar(100);unique_index"`
	FirstName string `json:"first_name"	form:"first_name"		query:"first_name"`
	LastName  string `json:"last_name" 	form:"last_name"		query:"last_name"`
}

// Validate Validates the data
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid Email")
	}
	return nil
}
