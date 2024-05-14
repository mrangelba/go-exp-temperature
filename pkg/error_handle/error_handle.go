package error_handle

import (
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type RestError struct {
	status int
	err    error
}

func (e RestError) Error() string {
	return e.err.Error()
}

func (e RestError) Response() (int, any) {
	return e.status, ErrorResponse{
		Message: e.Error(),
	}
}

func Handle(err error) (int, any) {
	var restErr RestError
	if errors.As(err, &restErr) {
		return restErr.Response()
	}

	return http.StatusInternalServerError, ErrorResponse{
		Message: err.Error(),
	}
}

var (
	ErrNotFound            = &RestError{status: http.StatusNotFound, err: errors.New("can not found zipcode")}
	ErrUnprocessableEntity = &RestError{status: http.StatusUnprocessableEntity, err: errors.New("invalid zipcode")}
)
