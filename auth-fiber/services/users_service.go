package services

import (
	"github.com/adharshmk96/go-microservices/auth-fiber/models/mysql/userdata"
	"github.com/adharshmk96/go-microservices/auth-fiber/utils/errors"
)

type userServices struct{}

// userServicesInterface
type userServicesInterface interface {
	CreateUser(user userdata.User) (*userdata.User, *errors.RestErr)
	GetUser(userID uint64) (*userdata.User, *errors.RestErr)
	UpdateUser(isPartial bool, user userdata.User) (*userdata.User, *errors.RestErr)
	DeleteUser(userID uint64) *errors.RestErr
	FindUser(status string) ([]userdata.User, *errors.RestErr)
}

// usersService structure
// type usersService struct{}

var (
	// UserServices interface
	UserServices userServicesInterface = &userServices{}
)

// CreateUser returns validation errors
func (s *userServices) CreateUser(user userdata.User) (*userdata.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Create(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser service returls user details
func (s *userServices) GetUser(userID uint64) (*userdata.User, *errors.RestErr) {
	result := &userdata.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateUser service updates user details
func (s *userServices) UpdateUser(isPartial bool, user userdata.User) (*userdata.User, *errors.RestErr) {

	current := &userdata.User{ID: user.ID}
	if err := current.Get(); err != nil {
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

// DeleteUser Deletes a user
func (s *userServices) DeleteUser(userID uint64) *errors.RestErr {
	user := &userdata.User{ID: userID}
	if err := user.Get(); err != nil {
		return err
	}
	return user.Delete()
}

// FindUser finds user by status
func (s *userServices) FindUser(status string) ([]userdata.User, *errors.RestErr) {
	user := &userdata.User{}
	return user.Find(status)
}
