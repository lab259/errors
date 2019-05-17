package gqerrors

import (
	"fmt"
	"strings"

	"github.com/lab259/errors"
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
			return false, errors.New(fmt.Sprintf("couldn't have key `module` \n%q", graphQLError))
		}

		switch matcher.Module.(type) {
		case errors.Option:
			option := matcher.Module.(errors.Option)
			mModule := option(nil).Error()
			return matcher.assert(mModule, module)
		case string:
			expected := matcher.Module.(string)
			if module != expected {
				return false, errors.New(fmt.Sprintf("expected module [%s] not equal [%s]", module, expected))
			}
		}
	}

	return true, nil
}

func (matcher *errWithGraphQLModuleMatcher) FailureMessage(actual interface{}) string {
	return format.Message(actual, "to have any module equal field [", matcher.Module, "]")
}

func (matcher *errWithGraphQLModuleMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(actual, "to have any module equal field [", matcher.Module, "]")
}

func ErrWithGraphQLModule(mutateOrQueryName string, module interface{}) *errWithGraphQLModuleMatcher {
	return &errWithGraphQLModuleMatcher{
		Module:        module,
		MutateOrQuery: mutateOrQueryName,
	}
}

func (matcher *errWithGraphQLModuleMatcher) assert(source, expected string) (bool, error) {
	if ok := strings.Contains(source, expected); !ok {
		return false, errors.New(fmt.Sprintf("expected module [%s] not equal [%s]", expected, source))
	}
	return true, nil
}
