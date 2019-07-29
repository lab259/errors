package gqlerrors

import (
	"fmt"

	"github.com/lab259/errors/v2"
	"github.com/onsi/gomega/format"
)

type errWithGraphQLCodeMatcher struct {
	Code          interface{}
	MutateOrQuery string
}

func (matcher *errWithGraphQLCodeMatcher) Match(actual interface{}) (bool, error) {
	graphQLError, err := prepareFromJSON(matcher.MutateOrQuery, actual)
	if err != nil {
		return false, err
	}

	for _, v := range graphQLError.Errors {
		code, ok := v.Extensions["code"].(string)
		if !ok {
			return false, fmt.Errorf("couldn't have key `code` %q", v.Extensions)
		}

		switch matcher.Code.(type) {
		case errors.Option:
			c := matcher.Code.(errors.Option)
			cError := c(nil).Error()
			if code != cError {
				return false, fmt.Errorf("code [%s] not equal [%s]", cError, code)

			}
		case string:
			if code != matcher.Code {
				return false, fmt.Errorf("code [%s] not equal [%s]", matcher.Code, code)
			}
		case nil:
			return false, fmt.Errorf("the code cannot be null")
		}

	}

	return true, nil
}

func (matcher *errWithGraphQLCodeMatcher) FailureMessage(actual interface{}) string {
	return format.Message(actual, "to have any code equal field", matcher.Code)
}

func (matcher *errWithGraphQLCodeMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(actual, "to have any code equal field", matcher.Code)
}

func ErrWithGraphQLCode(mutateOrQueryName string, code interface{}) *errWithGraphQLCodeMatcher {
	return &errWithGraphQLCodeMatcher{
		Code:          code,
		MutateOrQuery: mutateOrQueryName,
	}
}
