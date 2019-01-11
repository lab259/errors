package errors

// ErrorWithCode implements an error with a code related to it.
type ErrorWithCode interface {
	Error() string
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
	return err.code
}

// AppendData adds the code to the ErrorResponse.
func (err *errorWithCode) AppendData(response ErrorResponse) {
	response.SetParam("code", err.code)
}

// Reason is the error that originally was raised.
func (err *errorWithCode) Reason() error {
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
