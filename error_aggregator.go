package errors

// ErrorResponse is an abstraction of an error response that might, or not, be
// sent to the customer.
type ErrorResponse interface {
	SetParam(name string, value interface{})
}

// ErrorResponseAggregator is an error that can modify an `ErrorResponse`. It,
// usually, adds more information to explain better the error.
type ErrorResponseAggregator interface {
	AppendData(response ErrorResponse)
}
