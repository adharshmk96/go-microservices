package services

import (
	"github.com/adharshmk96/go-microservices/auth-gin/domain/users"
	cryptoutils "github.com/adharshmk96/go-microservices/auth-gin/utils/cryptoUtils"
	"github.com/adharshmk96/go-microservices/auth-gin/utils/dateutils"
	"github.com/adharshmk96/go-microservices/auth-gin/utils/errors"
)

type usersService struct{}

var (
	// UsersService interface
	UsersService userServiceInterface = &usersService{}
)

// userServiceInterface to access all functions
type userServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	FindUser(string) (users.Users, *errors.RestErr)
}

// CreateUser performs the business logic validation
func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = dateutils.GetDbString()
	user.Password = cryptoutils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil

}

// GetUser Used to retrieve a user
func (s *usersService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateUser used ot update a user's details
func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current := &users.User{ID: user.ID}
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

// DeleteUser deletes the user
func (s *usersService) DeleteUser(userID int64) *errors.RestErr {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return err
	}
	user := &users.User{ID: userID}
	return user.Delete()
}

// FindUser  finds by status
func (s *usersService) FindUser(status string) (users.Users, *errors.RestErr) {
	user := &users.User{}
	return user.Find(status)
}
