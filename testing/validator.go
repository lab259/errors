package testing

import (
	"errors"
	"strings"

	"github.com/onsi/gomega/format"
	"gopkg.in/go-playground/validator.v9"
)

type ErrorWithValidatorMatcher struct {
	Expected interface{}
}

func (matcher *ErrorWithValidatorMatcher) Match(actual interface{}) (r bool, rErr error) {
	err, ok := actual.(error)
	if !ok {
		return false, errors.New("`actual` is not an `error`")
	}
	if !reasonIterator(err, func(err error) bool {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			// returning false means that the iterator will continue going through the errors reasons.
			return false
		}

		e := appendMapErrors(validationErrors)

		// returning true means that the iterator is satisfied.
		r = checkFieldsMatcher(matcher, e)
		return checkFieldsMatcher(matcher, e)
	}) {
		return false, errors.New(format.Message(actual, "to not have any error equal", matcher.Expected))
	}

	return
}

func (matcher *ErrorWithValidatorMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to have any validation equal ", matcher.Expected)
}

func (matcher *ErrorWithValidatorMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to not have any validation equal", matcher.Expected)
}

var replacer = strings.NewReplacer("[", ".", "]", "")

func replaceBrackets(field string) string {
	return replacer.Replace(field)
}

func getFieldName(namespace string) string {
	parts := strings.SplitN(namespace, ".", 2)
	if len(parts) > 1 {
		return replaceBrackets(parts[1])
	}
	return replaceBrackets(namespace)
}

func appendMapErrors(validationErrors validator.ValidationErrors) map[string][]string {
	e := make(map[string][]string, 0)
	for _, validationErr := range validationErrors {
		field := getFieldName(validationErr.Namespace())
		e[field] = append(e[field], validationErr.ActualTag())
	}

	return e
}

func checkFieldsMatcher(matcher *ErrorWithValidatorMatcher, e map[string][]string) (result bool) {
	if expected, ok := matcher.Expected.([]string); ok {
		for _, k := range expected {
			if _, ok := e[k]; ok {
				result = true
			} else {
				result = false
			}
		}
	}

	return result
}

// ErrorWithValidation goes through the error chain verifying if there is any `Validation` is equal to errors
func ErrorWithValidation(errs ...string) *ErrorWithValidatorMatcher {
	return &ErrorWithValidatorMatcher{
		Expected: errs,
	}
}
