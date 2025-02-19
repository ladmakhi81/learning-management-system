package baseerror

import "net/http"

type ClientErr struct {
	Message    string `json:"message"`
	StatusCode uint   `json:"statusCode"`
}

type ClientValidationErr struct {
	ClientErr
	Errors any `json:"errors"`
}

type ServerErr struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

func (e ClientErr) Error() string {
	return e.Message
}

func (e ClientValidationErr) Error() string {
	return e.Message
}

func (e ServerErr) Error() string {
	return e.Message
}

func NewClientErr(message string, statusCode uint) ClientErr {
	return ClientErr{
		Message:    message,
		StatusCode: statusCode,
	}
}

func NewClientValidationErr(errors any) ClientValidationErr {
	return ClientValidationErr{
		Errors:    errors,
		ClientErr: NewClientErr("Invalid Input Body", http.StatusBadRequest),
	}
}

func NewServerErr(message, location string) ServerErr {
	return ServerErr{
		Message:  message,
		Location: location,
	}
}
