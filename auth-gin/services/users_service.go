package services

import (
	"github.com/adharshmk96/go-microservices/auth-gin/domain/users"
	"github.com/adharshmk96/go-microservices/auth-gin/utils/errors"
)

// CreateUser performs the business logic validation
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil

}

func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
