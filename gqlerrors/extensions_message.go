package gqlerrors

import (
	"fmt"
	"strings"

	"github.com/lab259/errors/v2"
	"github.com/onsi/gomega/format"
)

type errWithGraphQLMessageMatcher struct {
	Message       interface{}
	MutateOrQuery string
}

func (matcher *errWithGraphQLMessageMatcher) Match(actual interface{}) (bool, error) {
	graphQLError, err := prepareFromJSON(matcher.MutateOrQuery, actual)
	if err != nil {
		return false, err
	}

	for _, v := range graphQLError.Errors {
		if extMessage, ok := v.Extensions["message"].(string); ok {
			switch matcher.Message.(type) {
			case errors.Option:
				op := matcher.Message.(errors.Option)
				message := op(nil).Error()
				return matcher.messageOption(extMessage, message)
			case string:
				message := matcher.Message.(string)
				return matcher.messageOption(extMessage, message)
			}
		}

		if v.Message != "" {
			switch matcher.Message.(type) {
			case errors.Option:
				op := matcher.Message.(errors.Option)
				message := op(nil).Error()
				return matcher.messageOption(v.Message, message)
			case string:
				message := matcher.Message.(string)
				return matcher.messageOption(v.Message, message)
			case error:
				message := matcher.Message.(error)
				return matcher.messageOption(v.Message, message.Error())
			}
		}
	}

	return false, fmt.Errorf("the field message not found")
}

func (matcher *errWithGraphQLMessageMatcher) FailureMessage(actual interface{}) string {
	return format.Message(actual, "to have any message equal field", matcher.Message)
}

func (matcher *errWithGraphQLMessageMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(actual, "to have any message equal field", matcher.Message)
}

func ErrWithGraphQLMessage(mutateOrQueryName string, message interface{}) *errWithGraphQLMessageMatcher {
	return &errWithGraphQLMessageMatcher{
		Message:       message,
		MutateOrQuery: mutateOrQueryName,
	}
}

func (matcher *errWithGraphQLMessageMatcher) messageOption(source, expected string) (bool, error) {
	ok := strings.Contains(source, expected)
	if !ok {
		message := matcher.Message
		switch matcher.Message.(type) {
		case errors.Option:
			op := matcher.Message.(errors.Option)
			message = op(nil).Error()
		}
		return false, fmt.Errorf("message [%s] not equal [%s]", message, source)
	}
	return true, nil
}
