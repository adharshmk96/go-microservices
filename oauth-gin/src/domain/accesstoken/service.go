package accesstoken

import "github.com/adharshmk96/go-microservices/oauth-gin/src/utils/errors"

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{}
}

func (s *service) GetById(string) (*AccessToken, *errors.RestErr) {
	return nil, nil
}
