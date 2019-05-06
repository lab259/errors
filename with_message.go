package errors

import "fmt"

// ErrorWithMessage implements an error with a code related to it.
type ErrorWithMessage interface {
	Message() string
}

type errorWithMessage struct {
	reason  error
	message string
}

// Code is the code of the error.
//
// The error codes must be direct and have no repeated declarations.
func (err *errorWithMessage) Message() string {
	return err.message
}

// Error returns the code of the error
func (err *errorWithMessage) Error() string {
	if err.reason == nil {
		return err.message
	}
	return fmt.Sprintf("%s: %s", err.message, err.reason.Error())
}

// AppendData adds the code to the ErrorResponse.
func (err *errorWithMessage) AppendData(response ErrorResponse) {
	response.SetParam("message", err.message)
}

// Unwrap returns the next error in the error chain.
// If there is no next error, Unwrap returns nil.
func (err *errorWithMessage) Unwrap() error {
	return err.reason
}

// WrapMessage creates a new instance of a Reportable error with the given
// code.
func WrapMessage(reason error, code string) error {
	return &errorWithMessage{
		reason:  reason,
		message: code,
	}
}
