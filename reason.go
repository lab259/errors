package errors

// ErrorWithReason is an error that has an main reason because it was raised. It
// is used to wrap more abstract errors with improved messages for the customer.
type ErrorWithReason interface {
	Reason() error
}
