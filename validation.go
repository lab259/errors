package errors

import (
	"github.com/valyala/fasthttp"
	"gopkg.in/go-playground/validator.v9"
	"strings"
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
}

// Code returns the "validation" error code.
func (err *ValidationError) Code() string {
	return "validation"
}

func (err *ValidationError) AppendData(response ErrorResponse) {
	response.SetParam("code", err.Code())
	response.SetParam("statusCode", fasthttp.StatusBadRequest)
	if validationErrors, ok := err.reason.(validator.ValidationErrors); ok {
		errors := make(map[string][]string, 0)
		for _, validationErr := range validationErrors {
			field := getFieldName(validationErr.Namespace())
			errors[field] = append(errors[field], validationErr.ActualTag())
		}
		response.SetParam("errors", errors)
	}
}

func (err *ValidationError) Reason() error {
	return err.reason
}

func (err *ValidationError) Error() string {
	return "validation"
}

func WrapValidation(reason error) error {
	return &ValidationError{
		reason: reason,
	}
}
