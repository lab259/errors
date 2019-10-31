package gqlerrors

import (
	"fmt"
	"strings"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

// HaveValidation succeeds if actual is a GraphQL Error that have the
// passed-in field validation errors.
func HaveValidation(field string, rules ...string) types.GomegaMatcher {
	return &haveValidationMatcher{
		field: field,
		rules: rules,
	}
}

type haveValidationMatcher struct {
	err    interface{}
	errors map[string]interface{}
	field  string
	rules  []string
}

func (matcher *haveValidationMatcher) Match(actual interface{}) (ok bool, err error) {
	gqlerror, err := prepare("HaveValidation", actual)
	if err != nil {
		return false, err
	}

	var errors map[string]interface{}
	if gqlerror.Gqlerror != nil {
		errors, ok = gqlerror.Gqlerror.Extensions["errors"].(map[string]interface{})
		if !ok {
			return false, fmt.Errorf("Validation extension not found in %s", gqlerror)
		}

		matcher.err = gqlerror.Gqlerror
	}

	if gqlerror.FormattedError != nil {
		errors, ok = gqlerror.FormattedError.Extensions["errors"].(map[string]interface{})
		if !ok {
			return false, fmt.Errorf("Validation extension not found in %s", gqlerror)
		}

		matcher.err = gqlerror.FormattedError
	}

	if gqlerror.Gqlerror == nil && gqlerror.FormattedError == nil {
		return false, fmt.Errorf("Validation extension not found in %s", gqlerror)
	}

	matcher.errors = errors

	var checked bool
	if rules, ok := errors[matcher.field].([]interface{}); ok {
		for _, rule := range matcher.rules {
			var found bool

			for _, r := range rules {
				if r == rule {
					found = true
					break
				}
			}

			if !found {
				return false, nil
			}
		}

		checked = true
	}
	return checked, nil
}

func (matcher *haveValidationMatcher) FailureMessage(actual interface{}) string {
	return format.Message(matcher.err, fmt.Sprintf("to have Validation extension with failures (%s) for %s field", strings.Join(matcher.rules, ","), matcher.field))
}

func (matcher *haveValidationMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(matcher.err, fmt.Sprintf("to not have Validation extension with failures (%s) for %s field", strings.Join(matcher.rules, ","), matcher.field))
}
