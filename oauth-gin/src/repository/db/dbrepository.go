package db

import (
	"fmt"

	"github.com/adharshmk96/go-microservices/oauth-gin/src/clients/cassandra"
	"github.com/adharshmk96/go-microservices/oauth-gin/src/domain/accesstoken"
	"github.com/adharshmk96/go-microservices/oauth-gin/src/utils/errors"
)

func New() DBRepository {
	return &dbRepository{}
}

type DBRepository interface {
	GetById(string) (*accesstoken.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	_, err := cassandra.GetSession()
	if err != nil {
		panic("Error Creating a session")
	}
	fmt.Println("Im working ??")
	return nil, errors.NewBadRequestError("Not implemented")
}
