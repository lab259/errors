package errors_test

import (
	"errors"
	"log"
	"net/http"
	"testing"

	"github.com/jamillosantos/macchiato"
	lerrors "github.com/lab259/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var ErrTest = errors.New("this is a test")

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

	It("should wrap with a HTTPError (plain)", func() {
		nerr := errors.New("error")
		err := lerrors.Wrap(nerr, http.StatusNotFound)
		httpErr, ok := err.(lerrors.HttpError)
		Expect(ok).To(BeTrue())
		Expect(httpErr.StatusCode()).To(Equal(404))
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

	It("should wrap with a ErrorWithMessage (plain)", func() {
		nerr := errors.New("error")
		err := lerrors.Wrap(nerr, "123")
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

	It("should match wrapped error", func() {
		err := lerrors.Wrap(ErrTest, lerrors.Validation(), lerrors.Module("test"), lerrors.Message("message"))
		Expect(lerrors.Is(err, ErrTest)).To(BeTrue())
	})

	It("should return reason", func() {
		err := lerrors.Wrap(ErrTest, lerrors.Validation(), lerrors.Module("test"), lerrors.Message("message"))
		Expect(lerrors.Reason(err)).To(Equal(ErrTest))
	})

	It("should combine predefined options", func() {
		testModule := lerrors.Module("test")
		testOptionsCode := lerrors.Combine(lerrors.Code("test-options"), testModule)
		err := lerrors.Wrap(ErrTest, testOptionsCode)

		Expect(err.Error()).To(Equal("test: test-options: this is a test"))

		errModule, ok := err.(lerrors.ModuleError)
		Expect(ok).To(BeTrue())
		Expect(errModule.Module()).To(Equal("test"))

		errCode, ok := errModule.(lerrors.Wrapper).Unwrap().(lerrors.ErrorWithCode)
		Expect(ok).To(BeTrue())
		Expect(errCode.Code()).To(Equal("test-options"))

		errReason := errCode.(lerrors.Wrapper).Unwrap()
		Expect(errReason).To(Equal(ErrTest))
	})
})
