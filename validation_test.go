package errors_test

import (
	"errors"
	"reflect"

	ut "github.com/go-playground/universal-translator"
	lerrors "github.com/lab259/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/go-playground/validator.v9"
)

type fieldErrorMock struct {
	namespace string
}

func (*fieldErrorMock) Tag() string {
	return "Tag"
}

func (*fieldErrorMock) ActualTag() string {
	return "ActualTag"
}

func (ferr *fieldErrorMock) Namespace() string {
	return ferr.namespace
}

func (*fieldErrorMock) StructNamespace() string {
	return "StructNamespace"
}

func (*fieldErrorMock) Field() string {
	return "Field"
}

func (*fieldErrorMock) StructField() string {
	return "StructField"
}

func (*fieldErrorMock) Value() interface{} {
	return "Value"
}

func (*fieldErrorMock) Param() string {
	return "Param"
}

func (*fieldErrorMock) Kind() reflect.Kind {
	return reflect.String
}

func (*fieldErrorMock) Type() reflect.Type {
	return nil
}

func (*fieldErrorMock) Translate(ut ut.Translator) string {
	return "Translate"
}

var _ = Describe("ValidationError", func() {
	It("should build the errors in the http.ResponseError (wrapped+custom message)", func() {
		nerr := validator.ValidationErrors{
			&fieldErrorMock{
				namespace: "namespace",
			},
		}
		err := lerrors.Wrap(nerr, "custom message", lerrors.Validation(), lerrors.Code("validation-test"), lerrors.Module("test"))
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(`test: validation-test: custom message: "namespace" failed on [ActualTag]`))

		errResponse := NewMockErrorResponse()
		Expect(lerrors.AggregateToResponse(err, errResponse)).To(BeTrue())

		Expect(errResponse.Data).To(HaveKey("errors"))
		m, ok := errResponse.Data["errors"].(map[string][]string)
		Expect(ok).To(BeTrue())
		Expect(m).To(HaveKey("namespace"))
		Expect(m["namespace"]).To(ConsistOf("ActualTag"))
	})

	It("should build the errors in the http.ResponseError", func() {
		nerr := validator.ValidationErrors{
			&fieldErrorMock{
				namespace: "namespace",
			},
		}
		err := lerrors.WrapValidation(nerr)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(`validation: "namespace" failed on [ActualTag]`))

		errResponse := NewMockErrorResponse()

		validationErr, ok := err.(*lerrors.ValidationError)
		Expect(ok).To(BeTrue())
		validationErr.AppendData(errResponse)

		Expect(errResponse.Data).To(HaveKey("errors"))
		m, ok := errResponse.Data["errors"].(map[string][]string)
		Expect(ok).To(BeTrue())
		Expect(m).To(HaveKey("namespace"))
		Expect(m["namespace"]).To(ConsistOf("ActualTag"))
	})

	It("should build the errors in the http.ResponseError when namespace has dots", func() {
		nerr := validator.ValidationErrors{
			&fieldErrorMock{
				namespace: "namespace.fieldName",
			},
		}
		err := lerrors.WrapValidation(nerr)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(`validation: "fieldName" failed on [ActualTag]`))

		errResponse := NewMockErrorResponse()

		validationErr, ok := err.(*lerrors.ValidationError)
		Expect(ok).To(BeTrue())
		validationErr.AppendData(errResponse)

		reasonErr, ok := err.(lerrors.Wrapper)
		Expect(ok).To(BeTrue())
		Expect(reasonErr.Unwrap()).To(Equal(nerr))

		Expect(errResponse.Data).To(HaveKey("errors"))
		m, ok := errResponse.Data["errors"].(map[string][]string)
		Expect(ok).To(BeTrue())
		Expect(m).To(HaveKey("fieldName"))
		Expect(m["fieldName"]).To(ConsistOf("ActualTag"))
	})

	It("should not set errors when no ValidationErrors as reason", func() {
		nerr := errors.New("non validation error")
		err := lerrors.WrapValidation(nerr)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("validation"))

		errResponse := NewMockErrorResponse()

		validationErr, ok := err.(*lerrors.ValidationError)
		Expect(ok).To(BeTrue())
		validationErr.AppendData(errResponse)

		reasonErr, ok := err.(lerrors.Wrapper)
		Expect(ok).To(BeTrue())
		Expect(reasonErr.Unwrap()).To(Equal(nerr))

		Expect(errResponse.Data).ToNot(HaveKey("errors"))
	})
})
