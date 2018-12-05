package errors_test

import (
	lerrors "github.com/lab259/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

var _ = Describe("RecoverFromCrash", func() {
	It("should compose an error in a mock response", func() {
		nerr := errors.New("test")
		err := lerrors.Wrap(nerr, lerrors.Http(123), lerrors.Code("code1"), lerrors.Module("module1"))
		response := NewMockErrorResponse()
		func() {
			defer lerrors.RecoveryFromCrash(response, func(data interface{}) {
				Fail("this should not be called")
			})
			panic(err)
		}()
		Expect(response.Data).To(HaveKeyWithValue("statusCode", 123))
		Expect(response.Data).To(HaveKeyWithValue("code", "code1"))
		Expect(response.Data).To(HaveKeyWithValue("module", "module1"))
	})

	It("should call the default unknown handler for unhandled errors", func() {
		nerr := errors.New("test")
		response := NewMockErrorResponse()
		handlerCalled := false
		func() {
			defer lerrors.RecoveryFromCrash(response, func(data interface{}) {
				handlerCalled = true
				Expect(data).To(Equal(nerr))
			})
			panic(nerr)
		}()
		Expect(handlerCalled).To(BeTrue())
	})

	It("should call the default unknown handler for non error recovery data", func() {
		response := NewMockErrorResponse()
		handlerCalled := false
		func() {
			defer lerrors.RecoveryFromCrash(response, func(data interface{}) {
				handlerCalled = true
				Expect(data).To(Equal("not an error"))
			})
			panic("not an error")
		}()
		Expect(handlerCalled).To(BeTrue())
	})
})
