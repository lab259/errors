package testing

import (
	"github.com/lab259/errors"
	"github.com/onsi/gomega/format"
)

// ErrorWithValidatorMatcher as struct for validating errors
type ErrorWithValidatorMatcher struct {
	Field string
	Rule  string
}

// Match ...
func (matcher *ErrorWithValidatorMatcher) Match(actual interface{}) (bool, error) {
	err, ok := actual.(error)
	if !ok {
		return false, errors.New("`actual` is not an `error`")
	}

	isMatcher := errorWithReasonIterator(err, func(err error) bool {
		errWithValidation, ok := err.(errors.ErrorWithValidation)
		if !ok {
			// returning false means that the iterator will continue going through the errors reasons.
			return false
		}

		// returning true means that the iterator is satisfied.
		return checkFieldsMatcher(matcher, errWithValidation.Errors())
	})

	if isMatcher {
		return true, nil
	}

	return false, errors.New(format.Message(actual, "to not have any error equal field[", matcher.Field, "] and rule [", matcher.Rule, "]"))

}

// FailureMessage ...
func (matcher *ErrorWithValidatorMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to have any validation equal ", matcher.Field)
}

// NegatedFailureMessage ...
func (matcher *ErrorWithValidatorMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to not have any validation equal", matcher.Field)
}

func checkFieldsMatcher(matcher *ErrorWithValidatorMatcher, e map[string][]string) bool {
	if rules, ok := e[matcher.Field]; ok {
		for _, rule := range rules {
			if matcher.Rule == rule {
				return true
			}
		}
	}
	return false
}

// ErrorWithValidation goes through the error chain verifying if there is any `Validation` is equal to errors
func ErrorWithValidation(field string, rule string) *ErrorWithValidatorMatcher {
	return &ErrorWithValidatorMatcher{
		Field: field,
		Rule:  rule,
	}
}
