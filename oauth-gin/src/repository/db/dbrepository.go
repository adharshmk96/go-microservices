package db

import (
	"fmt"

	"github.com/adharshmk96/go-microservices/oauth-gin/src/clients/cassandra"
	"github.com/adharshmk96/go-microservices/oauth-gin/src/domain/accesstoken"
	"github.com/adharshmk96/go-microservices/oauth-gin/src/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE  access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?,?,?,?)"
	queryUpdateAccessToken = "UPDATE access_tokens SET expires=? WHERE access_token=?"
)

type DBRepository interface {
	GetById(string) (*accesstoken.AccessToken, *errors.RestErr)
	Create(accesstoken.AccessToken) *errors.RestErr
	UpdateExpiry(accesstoken.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func NewRepository() DBRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	// session, err := cassandra.GetSession()
	// if err != nil {
	// 	return nil, errors.NewInternalServerError("Error Creating Session")
	// }

	var result accesstoken.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NotFoundError("Token Not Found")
		}
		return nil, errors.NewInternalServerError("Error Retrieving token")
	}

	fmt.Println("Im working ??")
	return &result, errors.NewBadRequestError("Not implemented")
}

func (r *dbRepository) Create(at accesstoken.AccessToken) *errors.RestErr {
	// session, err := cassandra.GetSession()
	// if err != nil {
	// 	return errors.NewInternalServerError("Error Creating Session")
	// }

	if err := cassandra.GetSession().Query(
		queryCreateAccessToken,
		at.AccessToken,
		at.UserID,
		at.ClientID,
		at.Expires,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil

}

func (r *dbRepository) UpdateExpiry(at accesstoken.AccessToken) *errors.RestErr {
	// session, err := cassandra.GetSession()
	// if err != nil {
	// 	return errors.NewInternalServerError("Error Creating Session")
	// }

	if err := cassandra.GetSession().Query(
		queryUpdateAccessToken,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil

}
