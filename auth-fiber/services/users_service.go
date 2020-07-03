package services

import (
	"github.com/adharshmk96/go-microservices/auth-fiber/models/mysql/userdata"
	"github.com/adharshmk96/go-microservices/auth-fiber/utils/errors"
)

// CreateUser returns validation errors
func CreateUser(user userdata.User) (*userdata.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Create(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser service returls user details
func GetUser(userID uint64) (*userdata.User, *errors.RestErr) {
	result := &userdata.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateUser service updates user details
func UpdateUser(isPartial bool, user userdata.User) (*userdata.User, *errors.RestErr) {

	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil

}
