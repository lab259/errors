package gqerrors

import (
	"fmt"

	"github.com/lab259/errors"
	"github.com/onsi/gomega/format"
)

type errWithGraphQLValidatorMatcher struct {
	Field         string
	Rule          string
	MutateOrQuery string
}

func (matcher *errWithGraphQLValidatorMatcher) Match(actual interface{}) (bool, error) {
	graphQLError, err := prepare(matcher.MutateOrQuery, actual)
	if err != nil {
		return false, err
	}

	for _, v := range graphQLError.Errors {
		validation, ok := v.Extensions["errors"].(map[string]interface{})
		if !ok {
			return false, errors.New(fmt.Sprintf("couldn't have key `validation` on the %q", graphQLError))
		}

		// Matcher field
		if !checkFieldsMatcher(matcher, validation) {
			return false, errors.New(fmt.Sprintf("errors not containing key [%s] on the %s", matcher.Field, graphQLError.Errors))
		}
	}

	return true, nil
}

func (matcher *errWithGraphQLValidatorMatcher) FailureMessage(actual interface{}) string {
	return format.Message(actual, "to have any validation equal field [", matcher.Field, "] and rules [", matcher.Rule, "]")
}

func (matcher *errWithGraphQLValidatorMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(actual, "to have any validation equal field [", matcher.Field, "] and rules [", matcher.Rule, "]")
}

func ErrWithGraphQLValidate(mutateOrQueryName string, field string, rule string) *errWithGraphQLValidatorMatcher {
	return &errWithGraphQLValidatorMatcher{
		Field:         field,
		Rule:          rule,
		MutateOrQuery: mutateOrQueryName,
	}
}

func checkFieldsMatcher(matcher *errWithGraphQLValidatorMatcher, e map[string]interface{}) bool {
	if rules, ok := e[matcher.Field].([]interface{}); ok {
		for _, rule := range rules {
			if matcher.Rule == rule {
				return true
			}
		}
	}
	return false
}
