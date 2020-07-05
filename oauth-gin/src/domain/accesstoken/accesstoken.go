package accesstoken

import (
	"time"
)

const (
	expirationTime = 24
)

// AccessToken structure to send to client
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    string `json:"client_id"`
	Expires     int64  `json:"expires"`
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
