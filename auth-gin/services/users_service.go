package services

import (
	"net/http"

	"github.com/adharshmk96/go-microservices/auth-gin/domain/users"
	"github.com/adharshmk96/go-microservices/auth-gin/utils/errors"
)

// CreateUser performs the business logic validation
func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	return &user, nil

	return nil, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}
