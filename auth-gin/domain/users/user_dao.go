package users

// Retrieve and save -> database interaction

import (
	"fmt"
	"strings"

	"github.com/adharshmk96/go-microservices/auth-gin/datasources/mysql/userdatabase"
	"github.com/adharshmk96/go-microservices/auth-gin/utils/dateutils"
	"github.com/adharshmk96/go-microservices/auth-gin/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?)"
)

var (
	usersDB = make(map[int64]*User)
)

// Get gets user info
func (user *User) Get() *errors.RestErr {
	if err := userdatabase.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[user.ID]
	if result == nil {
		return errors.NewBadRequestError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

// Save Saves in db
func (user *User) Save() *errors.RestErr {
	stmt, err := userdatabase.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = dateutils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), "email_UNIQUE") {
			return errors.NewBadRequestError("Email already Exists")
		}

		return errors.NewInternalServerError(
			fmt.Sprintf("Erro adding user %s", err.Error()),
		)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("Error Trying to save user %s", err.Error()),
		)
	}

	user.ID = userID
	return nil

}
