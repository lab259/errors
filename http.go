package errors

type HttpError interface {
	StatusCode() int
}

// HttpError implements the aggregator to set the statusCode to an ErrorResponse.
type httpError struct {
	reason     error
	statusCode int
}

// Error returns the error message of the original error.
func (err *httpError) Error() string {
	return err.reason.Error()
}

// StatusCode returns the statusCode of the error.
func (err *httpError) StatusCode() int {
	return err.statusCode
}

// Unwrap returns the next error in the error chain.
// If there is no next error, Unwrap returns nil.
func (err *httpError) Unwrap() error {
	return err.reason
}

// AppendData aggregates the statusCode to the ErrorResponse.
func (err *httpError) AppendData(response ErrorResponse) {
	response.SetParam("statusCode", err.StatusCode())
}

// WrapHttp wraps a reason with an HttpError.
func WrapHttp(reason error, status int) error {
	return &httpError{
		reason:     reason,
		statusCode: status,
	}
}
