package errors

// ReportableError implements the idea of having an error that will be reported
// to the entity that is consuming the application.
type ReportableError interface {
	Code() string
}

type reportableError struct {
	reason error
	code   string
}

// Code is the code of the error.
//
// The error codes must be direct and have no repeated declarations.
func (err *reportableError) Code() string {
	return err.code
}

// Error returns the code of the error
func (err *reportableError) Error() string {
	return err.code
}

// AppendData adds the code to the ErrorResponse.
func (err *reportableError) AppendData(response ErrorResponse) {
	response.SetParam("code", err.code)
}

// Reason is the error that originally was raised.
func (err *reportableError) Reason() error {
	return err.reason
}

// NewReportable creates a new instance of a Reportable error with the given
// code.
func NewReportable(reason error, code string) error {
	return &reportableError{
		reason: reason,
		code:   code,
	}
}
