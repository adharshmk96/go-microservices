package atservice

import (
	"fmt"

	"github.com/adharshmk96/go-microservices/oauth-gin/src/domain/accesstoken"
	"github.com/adharshmk96/go-microservices/oauth-gin/src/repository/db"
	"github.com/adharshmk96/go-microservices/oauth-gin/src/utils/errors"
)

// type Repository interface {
// 	GetById(string) (*AccessToken, *errors.RestErr)
// }

type Service interface {
	GetById(string) (*accesstoken.AccessToken, *errors.RestErr)
	Create(accesstoken.AccessToken) *errors.RestErr
	UpdateExpiry(accesstoken.AccessToken) *errors.RestErr
}

type service struct {
	repository db.DBRepository
}

func NewService(repo db.DBRepository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*accesstoken.AccessToken, *errors.RestErr) {
	fmt.Println("Inside Service")
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(at accesstoken.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.Create(at)
}

func (s *service) UpdateExpiry(at accesstoken.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpiry(at)
}
