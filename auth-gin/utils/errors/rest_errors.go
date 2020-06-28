package errors

import "net/http"

// RestErr Generic structure for Rest errors
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// NewBadRequestError to return Bad request Error 400
func NewBadRequestError(messsage string) *RestErr {
	return &RestErr{
		Message: messsage,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}
