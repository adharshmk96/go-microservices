package accesstoken

// AccessToken structure to send to client
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    string `json:"client_id"`
}
