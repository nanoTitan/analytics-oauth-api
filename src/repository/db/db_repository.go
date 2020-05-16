package db

import (
	"log"

	"github.com/gocql/gocql"

	"github.com/nanoTitan/analytics-oauth-api/src/clients/cassandra"
	"github.com/nanoTitan/analytics-oauth-api/src/domain/atdomain"
	"github.com/nanoTitan/analytics-oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_token; SET expires=? WHERE access_token=?;"
)

// New - return a new Repository object
func New() Repository {
	return &dbRepository{}
}

// Repository - interface to interact with dbRepository objects
type Repository interface {
	GetByID(string) (*atdomain.AccessToken, *errors.RestErr)
	Create(atdomain.AccessToken) *errors.RestErr
	UpdateExpirationTime(atdomain.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

// GetById - Get an access token from the db given an Id
func (r *dbRepository) GetByID(id string) (*atdomain.AccessToken, *errors.RestErr) {
	var result atdomain.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires,
	); err != nil {

		if err.Error() == gocql.ErrNotFound.Error() {
			return nil, errors.NewNotFoundError("no access token found with given id")
		}
		log.Println(err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}

func (r *dbRepository) Create(at atdomain.AccessToken) *errors.RestErr {
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

func (r *dbRepository) UpdateExpirationTime(at atdomain.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(
		queryUpdateExpires,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
