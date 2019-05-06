package errors

import "fmt"

// ErrorWithCode implements an error with a code related to it.
type ErrorWithCode interface {
	Code() string
}

type errorWithCode struct {
	reason error
	code   string
}

// Code is the code of the error.
//
// The error codes must be direct and have no repeated declarations.
func (err *errorWithCode) Code() string {
	return err.code
}

// Error returns the code of the error
func (err *errorWithCode) Error() string {
	if err.reason == nil {
		return err.code
	}
	return fmt.Sprintf("%s: %s", err.code, err.reason.Error())
}

// AppendData adds the code to the ErrorResponse.
func (err *errorWithCode) AppendData(response ErrorResponse) {
	response.SetParam("code", err.code)
}

// Unwrap returns the next error in the error chain.
// If there is no next error, Unwrap returns nil.
func (err *errorWithCode) Unwrap() error {
	return err.reason
}

// WrapCode creates a new instance of a Reportable error with the given
// code.
func WrapCode(reason error, code string) error {
	return &errorWithCode{
		reason: reason,
		code:   code,
	}
}
