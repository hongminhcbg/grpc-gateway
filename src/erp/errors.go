package erp

import "fmt"

type ApplicationError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewApplicationError(code int, message string) *ApplicationError {
	return &ApplicationError{
		Code:    code,
		Message: message,
	}
}

func (e *ApplicationError) SetMessage(msg string) *ApplicationError {
	e.Message = msg
	return e
}

func (e *ApplicationError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}
