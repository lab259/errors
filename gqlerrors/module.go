package gqlerrors

import (
	"fmt"
	"reflect"

	"github.com/lab259/errors/v2"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
	"github.com/vektah/gqlparser/gqlerror"
)

// HaveModule succeeds if actual is a GraphQL Error that have the
// passed-in module extension.
func HaveModule(expected interface{}) types.GomegaMatcher {
	return &haveModuleMatcher{
		expected: expected,
	}
}

type haveModuleMatcher struct {
	err      gqlerror.Error
	module   string
	expected interface{}
}

func (matcher *haveModuleMatcher) Match(actual interface{}) (bool, error) {
	gqlerror, err := prepare("HaveModule", actual)
	if err != nil {
		return false, err
	}

	matcher.err = *gqlerror

	module, ok := gqlerror.Extensions["module"].(string)
	if !ok {
		return false, fmt.Errorf("Module extension not found in %s", gqlerror)
	}

	switch t := matcher.expected.(type) {
	case errors.ModuleError:
		matcher.module = t.Module()
	case errors.Option:
		if wrapErr, ok := t(nil).(errors.ModuleError); ok {
			matcher.module = wrapErr.Module()
		} else {
			return false, fmt.Errorf("HaveModule matcher only support an `errors.Module` option")
		}
	case string:
		matcher.module = t
	default:
		return false, fmt.Errorf("HaveModule matcher does not know how to assert %s", reflect.TypeOf(t))
	}

	return matcher.module == module, nil
}

func (matcher *haveModuleMatcher) FailureMessage(actual interface{}) string {
	return format.Message(matcher.err, "to have Module extension equals to", matcher.module)
}

func (matcher *haveModuleMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(matcher.err, "to not have Module extension equals to", matcher.module)
}
