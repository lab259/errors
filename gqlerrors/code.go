package gqlerrors

import (
	"fmt"
	"reflect"

	"github.com/lab259/errors/v2"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
	"github.com/vektah/gqlparser/gqlerror"
)

// HaveCode succeeds if actual is a GraphQL Error that have the
// passed-in code extension.
func HaveCode(expected interface{}) types.GomegaMatcher {
	return &haveCodeMatcher{
		expected: expected,
	}
}

type haveCodeMatcher struct {
	err      gqlerror.Error
	code     string
	expected interface{}
}

func (matcher *haveCodeMatcher) Match(actual interface{}) (bool, error) {
	gqlerror, err := prepare("HaveCode", actual)
	if err != nil {
		return false, err
	}

	matcher.err = *gqlerror

	code, ok := gqlerror.Extensions["code"].(string)
	if !ok {
		return false, fmt.Errorf("Code extension not found in %s", gqlerror)
	}

	switch t := matcher.expected.(type) {
	case errors.ErrorWithCode:
		matcher.code = t.Code()
	case errors.Option:
		if wrapErr, ok := t(nil).(errors.ErrorWithCode); ok {
			matcher.code = wrapErr.Code()
		} else {
			return false, fmt.Errorf("HaveCode matcher only support an `errors.Code` option")
		}
	case string:
		matcher.code = t
	default:
		return false, fmt.Errorf("HaveCode matcher does not know how to assert %s", reflect.TypeOf(t))
	}

	return matcher.code == code, nil
}

func (matcher *haveCodeMatcher) FailureMessage(actual interface{}) string {
	return format.Message(matcher.err, "to have Code extension equals to", matcher.code)
}

func (matcher *haveCodeMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(matcher.err, "to not have Code extension equals to", matcher.code)
}
