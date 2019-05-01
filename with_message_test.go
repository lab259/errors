package errors_test

import (
	lerrors "github.com/lab259/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

var _ = Describe("ErrorWithMessage", func() {
	It("should wrap with an ErrorWithMessage", func() {
		nerr := errors.New("test")
		err := lerrors.WrapMessage(nerr, "error code")
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("error code: test"))

		errWithMessage, ok := err.(lerrors.ErrorWithMessage)
		Expect(ok).To(BeTrue())
		Expect(errWithMessage.Message()).To(Equal("error code"))

		errResponse := NewMockErrorResponse()

		aggErr, ok := err.(lerrors.ErrorResponseAggregator)
		Expect(ok).To(BeTrue())
		aggErr.AppendData(errResponse)
		Expect(errResponse.Data).To(HaveKeyWithValue("message", "error code"))

		errWithReason, ok := err.(lerrors.Wrapper)
		Expect(ok).To(BeTrue())
		Expect(errWithReason.Unwrap()).To(Equal(nerr))
	})
})
