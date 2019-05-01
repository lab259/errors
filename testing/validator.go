package testing

import (
	"errors"
	"strings"

	"github.com/onsi/gomega/format"
	"gopkg.in/go-playground/validator.v9"
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

	isMatcher := reasonIterator(err, func(err error) bool {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			// returning false means that the iterator will continue going through the errors reasons.
			return false
		}

		e := appendMapErrors(validationErrors)

		// returning true means that the iterator is satisfied.
		return checkFieldsMatcher(matcher, e)
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
