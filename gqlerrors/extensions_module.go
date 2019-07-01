package gqlerrors

import (
	"fmt"
	"strings"

	"github.com/lab259/errors/v2"
	"github.com/onsi/gomega/format"
)

type errWithGraphQLModuleMatcher struct {
	Module        interface{}
	MutateOrQuery string
}

func (matcher *errWithGraphQLModuleMatcher) Match(actual interface{}) (bool, error) {
	graphQLError, err := prepare(matcher.MutateOrQuery, actual)
	if err != nil {
		return false, err
	}

	for _, v := range graphQLError.Errors {
		module, ok := v.Extensions["module"].(string)
		if !ok {
			return false, fmt.Errorf("couldn't have key `module` %v", graphQLError)
		}

		switch matcher.Module.(type) {
		case errors.Option:
			option := matcher.Module.(errors.Option)
			mModule := option(nil).Error()
			if ok := strings.Contains(mModule, module); !ok {
				return false, fmt.Errorf("expected module [%s] not equal [%s]", mModule, module)
			}
		case string:
			expected := matcher.Module.(string)
			if module != expected {
				return false, fmt.Errorf("expected module [%s] not equal [%s]", module, expected)
			}
		}
	}

	return true, nil
}

func (matcher *errWithGraphQLModuleMatcher) FailureMessage(actual interface{}) string {
	return format.Message(actual, "to have any module equal field", matcher.Module)
}

func (matcher *errWithGraphQLModuleMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(actual, "to have any module equal field", matcher.Module)
}

func ErrWithGraphQLModule(mutateOrQueryName string, module interface{}) *errWithGraphQLModuleMatcher {
	return &errWithGraphQLModuleMatcher{
		Module:        module,
		MutateOrQuery: mutateOrQueryName,
	}
}
