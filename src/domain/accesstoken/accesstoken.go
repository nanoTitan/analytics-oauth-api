package accesstoken

import (
	"strings"
	"time"

	"github.com/nanoTitan/analytics-oauth-api/src/utils/errors"
)

const (
	expirationTime = 24
)

// AccessToken - An object representing an auth0 access token
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
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
func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired - returns true if this access token is expired or false otherwise
func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
