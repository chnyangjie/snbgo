package snbgo

import "fmt"

type APIError struct {
	Code    string
	Message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("snbgo: API error %s: %s", e.Code, e.Message)
}

var (
	ErrConfigInvalid = fmt.Errorf("snbgo: configuration is invalid")
	ErrTokenInvalid  = &APIError{Code: "002002", Message: "token invalid or expired"}
	ErrLoginNeeded   = &APIError{Code: "002003", Message: "login first"}
	ErrInvalidParam  = &APIError{Code: "002005", Message: "invalid parameter"}
)
