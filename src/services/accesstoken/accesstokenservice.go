package accesstoken

import (
	"strings"

	"github.com/nanoTitan/analytics-oauth-api/src/domain/atdomain"
	"github.com/nanoTitan/analytics-oauth-api/src/repository/db"
	"github.com/nanoTitan/analytics-oauth-api/src/repository/rest"

	"github.com/nanoTitan/analytics-oauth-api/src/utils/errors"
)

// Service - an interface for the access token service
type Service interface {
	GetByID(string) (*atdomain.AccessToken, *errors.RestErr)
	Create(atdomain.AccessTokenRequest) (*atdomain.AccessToken, *errors.RestErr)
	UpdateExpirationTime(atdomain.AccessToken) *errors.RestErr
}

type service struct {
	restUsersRepo rest.UsersRepository
	dbRepo        db.Repository
}

// NewService - returns a new auth service object
func NewService(usersRepo rest.UsersRepository, dbRepo db.Repository) Service {
	return &service{
		restUsersRepo: usersRepo,
		dbRepo:        dbRepo,
	}
}

// GetById - return an access token given an Id
func (s *service) GetByID(accessTokenID string) (*atdomain.AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.dbRepo.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request atdomain.AccessTokenRequest) (*atdomain.AccessToken, *errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// TODO: support both grant types - client_credentials and password

	user, err := s.restUsersRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	// Generate a new access token
	at := atdomain.GetNewAccessToken(user.Id)
	at.Generate()

	// Save the new access token to the DB
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}

	return &at, nil
}

func (s *service) UpdateExpirationTime(at atdomain.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}
