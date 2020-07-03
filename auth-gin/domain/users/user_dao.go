package users

// Retrieve and save -> database interaction

import (
	"fmt"

	"github.com/adharshmk96/go-microservices/auth-gin/datasources/mysql/userdatabase"
	"github.com/adharshmk96/go-microservices/auth-gin/utils/dateutils"
	"github.com/adharshmk96/go-microservices/auth-gin/utils/errors"
	"github.com/adharshmk96/go-microservices/auth-gin/utils/mysqlutils"
	"github.com/go-sql-driver/mysql"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?)"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?"
	queryUpdateUser  = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	queryDeleteUser  = "DELETE FROM users WHERE id=?"
)

var (
	usersDB = make(map[int64]*User)
)

// Get gets user info
func (user *User) Get() *errors.RestErr {
	stmt, err := userdatabase.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysqlutils.ParseError(getErr)
	}

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

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if saveErr != nil {
		sqlErr, ok := saveErr.(*mysql.MySQLError)
		if !ok {
			return errors.NewInternalServerError(fmt.Sprintf("Error when saving user: %s", saveErr.Error()))
		}

		switch sqlErr.Number {
		case 1062:
			return errors.NewInternalServerError("Email Already Exists")
		}
		return errors.NewInternalServerError("Email Already Exists")
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

// Update updates USER
func (user *User) Update() *errors.RestErr {
	stmt, err := userdatabase.Client.Prepare(queryUpdateUser)
	if err != nil {
		fmt.Println("Error from Here")
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return mysqlutils.ParseError(err)
	}
	return nil
}

// Delete remove user
func (user *User) Delete() *errors.RestErr {
	stmt, err := userdatabase.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		return mysqlutils.ParseError(err)
	}

	return nil
}
