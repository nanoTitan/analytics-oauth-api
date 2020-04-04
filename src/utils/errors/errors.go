package errors

import (
	"errors"
	"net/http"
)

// RestErr - A struct representing a REST error message
type RestErr struct {
	Message string
	Status  int
	Error   string
}

// NewError - return a basic New error given a message string
func NewError(msg string) error {
	return errors.New(msg)
}

// NewBadRequestError - return a 400 bad request error given a message string
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// NewNotFoundError - return a 404 not found error given a message string
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "bad_request",
	}
}

// NewInternalServerError - return a 500 internal server error given a message string
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

// NewDbError - return a custom 400 bad request error given a message string and error message
func NewDbError(message string, errMsg string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   errMsg,
	}
}
