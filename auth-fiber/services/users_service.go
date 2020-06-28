package services

import "github.com/adharshmk96/go-microservices/auth-fiber/domain/users"

// CreateUser returns validation errors
func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
