package gqlerrors

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/lab259/errors/v2"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
	"github.com/vektah/gqlparser/gqlerror"
)

// HaveMessage succeeds if actual is a GraphQL Error that have the
// passed-in message extension.
func HaveMessage(expected interface{}) types.GomegaMatcher {
	return &haveMessageMatcher{
		expected: expected,
	}
}

type haveMessageMatcher struct {
	err      gqlerror.Error
	message  string
	expected interface{}
}

func (matcher *haveMessageMatcher) Match(actual interface{}) (bool, error) {
	gqlerror, err := prepare("HaveMessage", actual)
	if err != nil {
		return false, err
	}

	matcher.err = *gqlerror

	message, ok := gqlerror.Extensions["message"].(string)
	if !ok {
		return false, fmt.Errorf("Message extension not found in %s", gqlerror)
	}

	switch t := matcher.expected.(type) {
	case errors.ErrorWithMessage:
		matcher.message = t.Message()
	case errors.Option:
		if wrapErr, ok := t(nil).(errors.ErrorWithMessage); ok {
			matcher.message = wrapErr.Message()
		} else {
			return false, fmt.Errorf("HaveMessage matcher only support an `errors.Message` option")
		}
	case string:
		matcher.message = t
	default:
		return false, fmt.Errorf("HaveMessage matcher does not know how to assert %s", reflect.TypeOf(t))
	}

	return matcher.message == message, nil
}

func (matcher *haveMessageMatcher) FailureMessage(actual interface{}) string {
	return format.Message(matcher.err, "to have Message extension equals to", matcher.message)
}

func (matcher *haveMessageMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(matcher.err, "to not have Message extension equals to", matcher.message)
}

// ContainMessage succeeds if actual is a GraphQL Error that contains the
// passed-in message.
func ContainMessage(expected interface{}) types.GomegaMatcher {
	return &containMessageMatcher{
		expected: expected,
	}
}

type containMessageMatcher struct {
	err      gqlerror.Error
	message  string
	expected interface{}
}

func (matcher *containMessageMatcher) Match(actual interface{}) (bool, error) {
	gqlerror, err := prepare("ContainMessage", actual)
	if err != nil {
		return false, err
	}

	matcher.err = *gqlerror

	message, ok := gqlerror.Extensions["message"].(string)
	if !ok {
		message = gqlerror.Message
	}

	switch t := matcher.expected.(type) {
	case errors.ErrorWithMessage:
		matcher.message = t.Message()
	case errors.Option:
		if wrapErr, ok := t(nil).(errors.ErrorWithMessage); ok {
			matcher.message = wrapErr.Message()
		} else {
			return false, fmt.Errorf("ContainMessage matcher only support an `errors.Message` option")
		}
	case string:
		matcher.message = t
	default:
		return false, fmt.Errorf("ContainMessage matcher does not know how to assert %s", reflect.TypeOf(t))
	}

	return strings.Contains(message, matcher.message), nil
}

func (matcher *containMessageMatcher) FailureMessage(actual interface{}) string {
	return format.Message(matcher.err, "to contain message", matcher.message)
}

func (matcher *containMessageMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(matcher.err, "not to contain message", matcher.message)
}
