package testing

import (
	"reflect"

	"github.com/onsi/gomega/format"

	"github.com/lab259/errors/v2"
)

// HttpErrorMatcher is the matcher that can find HttpError in the error chain and check its value.
type HttpErrorMatcher struct {
	Expected int
}

// Match iterates through the error chain trying to find the HttpError. If it is found, this method performs the
// the check.
func (matcher *HttpErrorMatcher) Match(actual interface{}) (r bool, rErr error) {
	err, ok := actual.(error)
	if !ok {
		return false, errors.New("`actual` is not an `error`")
	}
	if !errorWithReasonIterator(err, func(err error) bool {
		errWithCode, ok := err.(errors.HttpError)
		if !ok {
			// returning false means that the iterator will continue going through the errors reasons.
			return false
		}

		r, rErr = reflect.DeepEqual(errWithCode.StatusCode(), matcher.Expected), nil
		// returning true means that the iterator is satisfied.
		return true
	}) {
		return false, nil
	}

	return
}

// FailureMessage formats the error message for an error code not found.
func (matcher *HttpErrorMatcher) FailureMessage(actual interface{}) (message string) {
	err := actual.(error)
	if !errorWithReasonIterator(err, func(err error) bool {
		errWithCode, ok := err.(errors.HttpError)
		if !ok {
			// returning false means that the iterator will continue going through the errors reasons.
			return false
		}

		message = format.Message(errWithCode, "to have status code", matcher.Expected)
		// returning true means that the iterator is satisfied.
		return true
	}) {
		return format.Message(actual, "does not have HttpError", matcher.Expected)
	}

	return
}

// FailureMessage formats the error message for an error code not found.
func (matcher *HttpErrorMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	err := actual.(error)
	errorWithReasonIterator(err, func(err error) bool {
		errWithCode, ok := err.(errors.HttpError)
		if !ok {
			// returning false means that the iterator will continue going through the errors reasons.
			return false
		}

		message = format.Message(errWithCode, "not to have status code", matcher.Expected)
		// returning true means that the iterator is satisfied.
		return true
	})
	return
}

// HttpError goes through the error chain verifying if there is any `HttpError`. If an HttpError is found,
// it checks the Code. If it cannot find, an failure is reported.
func HttpStatus(code int) *HttpErrorMatcher {
	return &HttpErrorMatcher{
		Expected: code,
	}
}
