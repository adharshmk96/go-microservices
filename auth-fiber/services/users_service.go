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
