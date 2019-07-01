package testing

import (
	"reflect"

	"github.com/onsi/gomega/format"

	"github.com/lab259/errors/v2"
)

// ErrorWithCodeMatcher is the matcher that can find ErrorWithCode in the error chain and check its value.
type ErrorWithCodeMatcher struct {
	Expected interface{}
}

// Match iterates through the error chain trying to find the ErrorWithCode. If it is found, this method performs the
// the check.
func (matcher *ErrorWithCodeMatcher) Match(actual interface{}) (r bool, rErr error) {
	err, ok := actual.(error)
	if !ok {
		return false, errors.New("`actual` is not an `error`")
	}
	if !errorWithReasonIterator(err, func(err error) bool {
		errWithCode, ok := err.(errors.ErrorWithCode)
		if !ok {
			// returning false means that the iterator will continue going through the errors reasons.
			return false
		}

		r, rErr = reflect.DeepEqual(errWithCode.Code(), matcher.Expected), nil
		// returning true means that the iterator is satisfied.
		return true
	}) {
		return false, nil
	}

	return
}

// FailureMessage formats the error message for an error code not found.
func (matcher *ErrorWithCodeMatcher) FailureMessage(actual interface{}) (message string) {
	err := actual.(error)
	if !errorWithReasonIterator(err, func(err error) bool {
		errWithCode, ok := err.(errors.ErrorWithCode)
		if !ok {
			// returning false means that the iterator will continue going through the errors reasons.
			return false
		}

		message = format.Message(errWithCode, "to have code", matcher.Expected)
		// returning true means that the iterator is satisfied.
		return true
	}) {
		return format.Message(actual, "does not have ErrorWithCode", matcher.Expected)
	}

	return
}

// FailureMessage formats the error message for an error code not found.
func (matcher *ErrorWithCodeMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	err := actual.(error)
	errorWithReasonIterator(err, func(err error) bool {
		errWithCode, ok := err.(errors.ErrorWithCode)
		if !ok {
			// returning false means that the iterator will continue going through the errors reasons.
			return false
		}

		message = format.Message(errWithCode, "not to have code", matcher.Expected)
		// returning true means that the iterator is satisfied.
		return true
	})
	return
}

// ErrorWithCode goes through the error chain verifying if there is any `ErrorWithCode`. If an ErrorWithCode is found,
// it checks the Code. If it cannot find, an failure is reported.
func ErrorWithCode(code string) *ErrorWithCodeMatcher {
	return &ErrorWithCodeMatcher{
		Expected: code,
	}
}
