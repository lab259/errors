package errors

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/valyala/fasthttp"
	"gopkg.in/go-playground/validator.v9"
)

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

// ValidationError implements wrapping validation errors into a reportable
type ValidationError struct {
	reason error
	errors map[string][]string
}

// Code returns the "validation" error code.
func (err *ValidationError) Code() string {
	return "validation"
}

func (err *ValidationError) Errors() map[string][]string {
	if err.errors == nil {
		err.errors = make(map[string][]string, 0)
		if reason := Reason(err); reason != nil {
			if validationErrors, ok := reason.(validator.ValidationErrors); ok {
				for _, validationErr := range validationErrors {
					field := getFieldName(validationErr.Namespace())
					err.errors[field] = append(err.errors[field], validationErr.ActualTag())
				}
			}
		}
	}
	return err.errors
}

func (err *ValidationError) AppendData(response ErrorResponse) {
	response.SetParam("code", err.Code())
	response.SetParam("statusCode", fasthttp.StatusBadRequest)
	if errors := err.Errors(); len(errors) > 0 {
		response.SetParam("errors", errors)
	}
}

// Unwrap returns the next error in the error chain.
// If there is no next error, Unwrap returns nil.
func (err *ValidationError) Unwrap() error {
	return err.reason
}

func (err *ValidationError) Error() string {
	message := "validation"
	if err.reason == nil {
		return message
	}
	if errWithMessage, ok := err.reason.(ErrorWithMessage); ok {
		message = errWithMessage.Message()
	}
	errors := err.Errors()
	if len(errors) == 0 {
		return message
	}
	buff := bytes.NewBufferString("")
	for field, rules := range errors {
		if buff.Len() > 0 {
			buff.WriteString("; ")
		}
		buff.WriteString(fmt.Sprintf(`"%s" failed on %s`, field, rules))
	}
	return fmt.Sprintf("%s: %s", message, buff.String())
}

func WrapValidation(reason error) error {
	return &ValidationError{
		reason: reason,
	}
}
