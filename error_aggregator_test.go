package errors_test

import (
	"errors"
	lerrors "github.com/lab259/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RecoverFromCrash", func() {
	It("should compose an error in a mock response", func() {
		nerr := errors.New("test")
		err := lerrors.Wrap(nerr, lerrors.Http(123), lerrors.Code("code1"), lerrors.Module("module1"))
		response := NewMockErrorResponse()
		Expect(lerrors.AggregateToResponse(err, response)).To(BeTrue())
		Expect(response.Data).To(HaveKeyWithValue("statusCode", 123))
		Expect(response.Data).To(HaveKeyWithValue("code", "code1"))
		Expect(response.Data).To(HaveKeyWithValue("module", "module1"))
	})

	It("should call the default unknown handler for unhandled errors", func() {
		nerr := errors.New("test")
		response := NewMockErrorResponse()
		Expect(lerrors.AggregateToResponse(nerr, response)).To(BeFalse())
	})
})
