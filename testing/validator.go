package testing

import (
	"errors"
	errors2 "github.com/lab259/errors"
	"github.com/onsi/gomega/format"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

type ErrorWithValidatorMatcher struct {
	Expected interface{}
}

func (matcher *ErrorWithValidatorMatcher) Match(actual interface{}) (r bool, rErr error) {
	switch actual.(type) {
	case validator.ValidationErrors:
		validationErrors := actual.(validator.ValidationErrors)

		e := appendMapErrors(validationErrors)

		if checkFieldsMatcher(matcher, e) {
			return true, nil
		}

		return false, errors.New(format.Message(actual, "not to have any error equal", matcher.Expected))
	case *validator.ValidationErrors:
		if errs, ok := actual.(*validator.ValidationErrors); ok {

			e := appendMapErrors(*errs)

			if checkFieldsMatcher(matcher, e) {
				return true, nil
			}

		}

		return false, errors.New(format.Message(actual, "not to have any error equal", matcher.Expected))
	case errors2.ValidationError:
		errs := actual.(errors2.ValidationError)

		if validationErrors, ok := errs.Reason().(validator.ValidationErrors); ok {
			e := appendMapErrors(validationErrors)

			if checkFieldsMatcher(matcher, e) {
				return true, nil
			}

		}

		return false, errors.New(format.Message(actual, "not to have any error equal", matcher.Expected))
	case *errors2.ValidationError:
		errs := actual.(*errors2.ValidationError)

		if validationErrors, ok := errs.Reason().(validator.ValidationErrors); ok {
			e := appendMapErrors(validationErrors)

			if checkFieldsMatcher(matcher, e) {
				return true, nil
			}

		}

		return false, errors.New(format.Message(actual, "not to have any error equal", matcher.Expected))
	}

	return false, nil
}

func (matcher *ErrorWithValidatorMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to have any validation equal ", matcher.Expected)
}

func (matcher *ErrorWithValidatorMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to have any validation equal", matcher.Expected)
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
