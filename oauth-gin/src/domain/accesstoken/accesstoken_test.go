package accesstoken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "Expiry must be 24 hours")
}

// GetNewAccesToken to test retrieval of access token
func GetNewAccesToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "New token shouldn't be expired")
	assert.EqualValues(t, "", at.AccessToken, "New Token shouldn't have token id")
	assert True(t, at.UserID == 0,"New Access token shouldn't have associated user ID" )
}

// TestAccessTokenIsExpired
func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "Empty access token should be expired")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "Access token created three hours from now shouldn't be expired")
}
