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

// AggregateToResponse add the error information pieces to the errResponse. Then,
// it tries to check its `Reason` (if the error is an `ErrorWithReason`) for more
// data. It goes upward until `Reason` is not found.
//
// If at least one `ErrorResponseAggregator` is found it returns true, otherwise
// it returns false.
func AggregateToResponse(data interface{}, errResponse ErrorResponse) bool {
	if err, ok := data.(error); ok {
		handled := false
		for err != nil {
			if e, ok := err.(ErrorResponseAggregator); ok {
				e.AppendData(errResponse)
				handled = true
			}
			if e, ok := err.(Wrapper); ok {
				err = e.Unwrap()
				continue
			}
			break
		}
		if handled {
			return true
		}
	}
	return false
}
