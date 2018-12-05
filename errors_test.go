package errors_test

import (
	"errors"
	"github.com/jamillosantos/macchiato"
	lerrors "github.com/lab259/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"testing"
)

type MockErrorResponse struct {
	Data map[string]interface{}
}

func (response *MockErrorResponse) SetParam(name string, value interface{}) {
	response.Data[name] = value
}

func NewMockErrorResponse() *MockErrorResponse {
	return &MockErrorResponse{
		Data: make(map[string]interface{}),
	}
}

func TestErrors(t *testing.T) {
	log.SetOutput(GinkgoWriter)
	RegisterFailHandler(Fail)
	macchiato.RunSpecs(t, "Errors Test Suite")
}

var _ = Describe("Wrap", func() {
	It("should wrap with a HTTPError", func() {
		nerr := errors.New("error")
		err := lerrors.Wrap(nerr, lerrors.Http(123))
		httpErr, ok := err.(lerrors.HttpError)
		Expect(ok).To(BeTrue())
		Expect(httpErr.StatusCode()).To(Equal(123))
	})

	It("should wrap with a ErrorWithCode", func() {
		nerr := errors.New("error")
		err := lerrors.Wrap(nerr, lerrors.Code("123"))
		reportableErr, ok := err.(lerrors.ErrorWithCode)
		Expect(ok).To(BeTrue())
		Expect(reportableErr.Code()).To(Equal("123"))
	})

	It("should wrap with a ErrorWithMessage", func() {
		nerr := errors.New("error")
		err := lerrors.Wrap(nerr, lerrors.Message("123"))
		reportableErr, ok := err.(lerrors.ErrorWithMessage)
		Expect(ok).To(BeTrue())
		Expect(reportableErr.Message()).To(Equal("123"))
	})

	It("should wrap with a ModuleError", func() {
		nerr := errors.New("error")
		err := lerrors.Wrap(nerr, lerrors.Module("123"))
		moduleErr, ok := err.(lerrors.ModuleError)
		Expect(ok).To(BeTrue())
		Expect(moduleErr.Module()).To(Equal("123"))
	})

	It("should wrap with a ValidationError", func() {
		nerr := errors.New("error")
		err := lerrors.Wrap(nerr, lerrors.Validation())
		_, ok := err.(*lerrors.ValidationError)
		Expect(ok).To(BeTrue())
	})
})
