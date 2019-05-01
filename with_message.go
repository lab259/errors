package errors

// ErrorWithMessage implements an error with a code related to it.
type ErrorWithMessage interface {
	error
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
	return err.message
}

// AppendData adds the code to the ErrorResponse.
func (err *errorWithMessage) AppendData(response ErrorResponse) {
	response.SetParam("message", err.message)
}

// Reason is the error that originally was raised.
func (err *errorWithMessage) Unwrap() error {
	return err.reason
}

// WrapCode creates a new instance of a Reportable error with the given
// code.
func WrapMessage(reason error, code string) error {
	return &errorWithMessage{
		reason:  reason,
		message: code,
	}
}
