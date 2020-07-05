package users

import "encoding/json"

// PublicUser strucure for public
type PublicUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	// Password    string `json:"-"`
}

// PrivateUser strucure for public
type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	// Password    string `json:"-"`
}

// Users is group of users
type Users []User

// Marshall return public or private user list
func (users Users) Marshall(isPublic bool) interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}

// Marshall return public or private user struct
func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:          user.ID,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}

	userJSON, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJSON, &privateUser)
	return privateUser
}
