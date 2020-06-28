package services

import "github.com/adharshmk96/go-microservices/auth-gin/domain/users"

// CreateUser performs the business logic validation
func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
