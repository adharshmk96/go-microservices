package accesstoken

import (
	"strings"
	"time"

	"github.com/adharshmk96/go-microservices/oauth-gin/src/utils/errors"
)

const (
	expirationTime = 24
)

// AccessToken structure to send to client
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

// Validate to validate the access token
func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("Invalid token")
	}
	if at.UserID <= 0 {
		return errors.NewBadRequestError("Invalid User")
	}
	if at.ClientID <= 0 {
		return errors.NewBadRequestError("Invalid Client")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("Invalid Expiry")
	}
	return nil
}

// GetNewAccessToken returns Access Token
func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired method for access token to check if its expired
func (at *AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
