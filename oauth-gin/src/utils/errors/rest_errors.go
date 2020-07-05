package errors

import "net/http"

// RestErr Generic structure for Rest errors
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// NewBadRequestError to return Bad request Error 400
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// NotFoundError to return Bad request Error 404
func NotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "bad_request",
	}
}

// NewInternalServerError Returns 500
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
