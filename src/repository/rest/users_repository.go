package rest

import (
	"encoding/json"
	"time"

	"github.com/go-resty/resty"
	//_ "github.com/go-resty/resty/v2"
	"github.com/nanoTitan/analytics-oauth-api/src/utils/errors"
	"github.com/nanoTitan/analytics-users-api/domain/users"
)

var (
	usersRestClient = resty.New()
	baseURL         = "https://api.game_analytics.com"
	timeout         = 100 * time.Millisecond
)

// GetBaseURL - return the http base URL string
func GetBaseURL() string {
	return baseURL
}

// GetClient - returns the http users rest client
func GetClient() *resty.Client {
	return usersRestClient
}

// New - returns a new usersRepository object
func New() UsersRepository {
	return &usersRepository{}
}

// UsersRepository - An interface for users actions
type UsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct{}

// NewRepository - creates a new RestUsersRepository object
func NewRepository() UsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	requestBody := users.LoginRequest{
		Email:    email,
		Password: password,
	}

	resp, clientErr := usersRestClient.R().
		EnableTrace().
		SetBody(requestBody).
		Post(baseURL + "/users/login")

	if clientErr != nil {
		return nil, errors.NewInternalServerError(clientErr.Error())
	}

	if resp.StatusCode() > 299 {
		var restErr errors.RestErr
		respErr := json.Unmarshal(resp.Body(), &restErr)
		if respErr != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users response")
	}
	return &user, nil
}
