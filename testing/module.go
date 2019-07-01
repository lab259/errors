package testing

import (
	"reflect"

	"github.com/onsi/gomega/format"

	"github.com/lab259/errors/v2"
)

// ErrorWithModuleMatcher is the matcher that can find ErrorWithModule in the error chain and check its value.
type ErrorWithModuleMatcher struct {
	Expected interface{}
}

// Match iterates through the error chain trying to find the ErrorWithModule. If it is found, this method performs the
// the check.
func (matcher *ErrorWithModuleMatcher) Match(actual interface{}) (r bool, rErr error) {
	err, ok := actual.(error)
	if !ok {
		return false, errors.New("`actual` is not an `error`")
	}
	if !errorWithReasonIterator(err, func(err error) bool {
		errWithModule, ok := err.(errors.ModuleError)
		if !ok {
			// returning false means that the iterator will continue going through the errors reasons.
			return false
		}

		r, rErr = reflect.DeepEqual(errWithModule.Module(), matcher.Expected), nil
		// returning true means that the iterator is satisfied.
		return true
	}) {
		return false, nil
	}

	return
}

// FailureMessage formats the error message for an error module not found.
func (matcher *ErrorWithModuleMatcher) FailureMessage(actual interface{}) (message string) {
	err := actual.(error)
	if !errorWithReasonIterator(err, func(err error) bool {
		errWithModule, ok := err.(errors.ModuleError)
		if !ok {
			// returning false means that the iterator will continue going through the errors reasons.
			return false
		}

		message = format.Message(errWithModule.Module(), "to have module", matcher.Expected)
		// returning true means that the iterator is satisfied.
		return true
	}) {
		return format.Message(actual, "does not have ErrorWithModule", matcher.Expected)
	}

	return
}

// FailureMessage formats the error message for an error module not found.
func (matcher *ErrorWithModuleMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	err := actual.(error)
	errorWithReasonIterator(err, func(err error) bool {
		errWithModule, ok := err.(errors.ModuleError)
		if !ok {
			// returning false means that the iterator will continue going through the errors reasons.
			return false
		}

		message = format.Message(errWithModule.Module(), "not to have module", matcher.Expected)
		// returning true means that the iterator is satisfied.
		return true
	})
	return
}

// ErrorWithModule goes through the error chain verifying if there is any `ErrorWithModule`. If an ErrorWithModule is found,
// it checks the module. If it cannot find, an failure is reported.
func ErrorWithModule(module string) *ErrorWithModuleMatcher {
	return &ErrorWithModuleMatcher{
		Expected: module,
	}
}
