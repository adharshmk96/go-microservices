package mysqlutils

import (
	"strings"

	"github.com/adharshmk96/go-microservices/auth-gin/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

// ParseError is general mysql error handler
func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewInternalServerError("No matching records")
		}
		return errors.NewInternalServerError("Error parsing db response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("Invalid Data")
	}
	return errors.NewInternalServerError("Error processing request")

}
