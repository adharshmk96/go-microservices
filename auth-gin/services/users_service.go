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

// GetUser Used to retrieve a user
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateUser used ot update a user's details
func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
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

// DeleteUser deletes the user
func DeleteUser(userID int64) *errors.RestErr {
	if _, err := GetUser(userID); err != nil {
		return err
	}
	user := &users.User{ID: userID}
	return user.Delete()
}
