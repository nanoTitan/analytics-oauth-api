package atdomain

import (
	"fmt"
	"strings"
	"time"

	"github.com/nanoTitan/analytics-oauth-api/src/utils/crypto"
	"github.com/nanoTitan/analytics-oauth-api/src/utils/errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

// AccessTokenRequest - An object representing a grant-type access token
type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for client_credentials grant type
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// Validate - vaidates the fields of an AccessTokenRequest object
func (at *AccessTokenRequest) Validate() *errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break

	case grantTypeClientCredentials:
		break

	default:
		return errors.NewBadRequestError("invalid grant_type parameter")
	}

	// TODO: validate params for each grant_type
	return nil
}

// AccessToken - An object representing an auth0 access token
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id,omitempty"`
	Expires     int64  `json:"expires"`
}

// Validate - vaidates the fields of an AccessToken object
func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserID <= 0 {
		return errors.NewBadRequestError("invaid user id")
	}
	if at.ClientID <= 0 {
		return errors.NewBadRequestError("invaid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invaid expiration time")
	}

	return nil
}

// GetNewAccessToken - return a new access token with an expiration time
func GetNewAccessToken(userID int64) AccessToken {
	return AccessToken{
		UserID:  userID,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired - returns true if this access token is expired or false otherwise
func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

// Generate - creates an access token through a Md5 hash string
func (at *AccessToken) Generate() {
	at.AccessToken = crypto.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}
