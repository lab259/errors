package testing

import (
	"fmt"
	"reflect"

	"github.com/lab259/errors/v2"
	"github.com/onsi/gomega/format"
)

type ErrorWithReasonMatcher struct {
	Expected error
}

func (matcher *ErrorWithReasonMatcher) Match(actual interface{}) (r bool, rErr error) {
	if actual == nil && matcher.Expected == nil {
		return false, fmt.Errorf("Refusing to compare <nil> to <nil>.\nBe explicit and use BeNil() instead.  This is to avoid mistakes where both sides of an assertion are erroneously uninitialized.")
	}

	err, ok := actual.(error)
	if !ok {
		return false, errors.New("`actual` is not an `error`")
	}
	errorWithReasonIterator(err, func(err error) bool {
		if reflect.DeepEqual(err, matcher.Expected) {
			r, rErr = true, nil
			// returning true the iterator will stop
			return true
		}

		errorWithReason, ok := err.(errors.Wrapper)
		if !ok {
			// returning false means the iterator will stop iterating through the chain
			return false
		}

		if reflect.DeepEqual(errorWithReason.Unwrap(), matcher.Expected) {
			r, rErr = true, nil
			// returning true the iterator will stop
			return true
		}

		return false
	})

	return
}

func (matcher *ErrorWithReasonMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to have any reason equal", matcher.Expected)
}

func (matcher *ErrorWithReasonMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to have any reason equal", matcher.Expected)
}

// ErrorWithReason goes through the error chain verifying if there is any `Reason` is equal to the one
func ErrorWithReason(err error) *ErrorWithReasonMatcher {
	return &ErrorWithReasonMatcher{
		Expected: err,
	}
}
